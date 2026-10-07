// Command claim reads one claim out of the ledgers instead of the whole ledger.
//
// A ledger groups claims under topic headings. Each topic has an index table,
// one row per claim (ID, one-sentence headline, confidence grade, status,
// evidence), followed by one card per claim under a "### <ID>" heading with the
// facts, confidence reasoning, Unknowns and amendments. claims/registry.md
// defines the format. A ledger that has no cards yet is read row by row.
//
//	claim TERR-LOC-001 AI-DIPLO-004   one claim each, with its retraction entries
//	claim -k "sight range"            search every ledger, one line per hit
//	claim -l terrain                  one ledger's headlines as an index
//	claim -check                      lint every card ledger against the format
//	claim -stats                      what each ledger costs to open
//
// Exit 1 if any requested id is missing: an unknown id is a defect in the brief
// that asked for it, not an empty result.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

// A claim id is <FMT>-<TOPIC>-<NNN> or <FMT>-<NNN>. The topic segment may carry
// digits (SPR256, SPR16A), and the number may pass 999, so the pattern anchors
// on the trailing numeric suffix and leaves the middle segments optional. It is
// applied to a row's first cell or a card heading only, never to prose.
var idRe = regexp.MustCompile(`\b([A-Z][A-Z0-9]*(?:-[A-Z0-9]+)*-[0-9]{3,})\b`)

var (
	cardRe     = regexp.MustCompile(`^### ([A-Z][A-Z0-9]*(?:-[A-Z0-9]+)*-[0-9]{3,})\s*$`)
	headingRe  = regexp.MustCompile(`^#{1,6} `)
	fenceRe    = regexp.MustCompile("^\\s*(```|~~~)")
	gradeRe    = regexp.MustCompile(`^(High|Medium|Low|Unknown)( / (High|Medium|Low|Unknown))*$`)
	statusRe   = regexp.MustCompile(`^(✔ promoted|● active|✖ retracted)( \(([a-z ,]+)\))?$`)
	evidenceRe = regexp.MustCompile(`EXP-[0-9]{4}`)
	linkRe     = regexp.MustCompile(`\[([^\]]+)\]\([^)]*\)`)
)

// qualifiers is the closed vocabulary of a Status cell's parenthesis; true
// marks the ones that need an **Amended.** paragraph in the card.
var qualifiers = map[string]bool{
	"amended":             true,
	"partially retracted": true,
	"superseded":          true,
	"contested":           true,
	"branch candidate":    false,
}

const (
	headlineMax = 240  // characters in the Claim cell of a card ledger
	cardMax     = 8192 // bytes in one card
)

type row struct {
	id       string
	file     string // ledger base name, e.g. "terrain.md"
	line     int
	cells    []string
	raw      string
	card     string // card body without its heading; empty in a table-only ledger
	cardLine int
	hasCard  bool
}

// ledger is one parsed file: its rows in order, and cards that found no row.
type ledger struct {
	name    string
	rows    []row
	orphans []string // "<ID> at line N": a card whose ID has no row in this file
	dupes   []string // "<ID> at line N": a second card for the same ID
	inner   []string // "<ID> at line N": an id-keyed table row inside a card
	order   []string // card IDs in file order
}

func main() {
	var (
		keyword = flag.String("k", "", "search every ledger for a regexp and print one line per hit")
		list    = flag.String("l", "", "list one ledger's headlines as an index")
		stats   = flag.Bool("stats", false, "print what each ledger costs to open")
		check   = flag.Bool("check", false, "lint every card ledger against claims/registry.md")
		dir     = flag.String("dir", "", "claims directory (default: ./claims, then ../claims)")
		full    = flag.Bool("full", false, "with -k or -l, print whole claims instead of an index")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: claim [flags] ID [ID...]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	root, err := claimsDir(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "claim:", err)
		os.Exit(2)
	}
	ledgers, err := loadLedgers(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "claim:", err)
		os.Exit(2)
	}
	var rows []row
	for _, l := range ledgers {
		rows = append(rows, l.rows...)
	}

	switch {
	case *check:
		os.Exit(lint(ledgers))
	case *stats:
		printStats(root, rows)
	case *list != "":
		os.Exit(printLedger(rows, *list, *full))
	case *keyword != "":
		os.Exit(printSearch(rows, *keyword, *full))
	default:
		if flag.NArg() == 0 {
			flag.Usage()
			os.Exit(2)
		}
		os.Exit(printIDs(rows, flag.Args()))
	}
}

func claimsDir(override string) (string, error) {
	if override != "" {
		return override, nil
	}
	for _, c := range []string{"claims", filepath.Join("..", "claims"), filepath.Join("..", "..", "claims")} {
		if fi, err := os.Stat(c); err == nil && fi.IsDir() {
			return c, nil
		}
	}
	return "", fmt.Errorf("no claims/ directory found (run from the repo root, or pass -dir)")
}

// load returns every row of every ledger, cards attached.
func load(root string) ([]row, error) {
	ledgers, err := loadLedgers(root)
	if err != nil {
		return nil, err
	}
	var out []row
	for _, l := range ledgers {
		out = append(out, l.rows...)
	}
	return out, nil
}

func loadLedgers(root string) ([]ledger, error) {
	files, err := filepath.Glob(filepath.Join(root, "*.md"))
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no ledgers under %s", root)
	}
	sort.Strings(files)
	var out []ledger
	for _, f := range files {
		l, err := parseLedger(f)
		if err != nil {
			return nil, err
		}
		out = append(out, l)
	}
	return out, nil
}

// parseLedger reads one file. A row is a table line outside a card and outside
// a code fence whose first cell holds a claim id. A card runs from its "### ID"
// heading to the next heading; table lines inside it belong to the card.
func parseLedger(path string) (ledger, error) {
	l := ledger{name: filepath.Base(path)}
	fh, err := os.Open(path)
	if err != nil {
		return l, err
	}
	defer fh.Close()

	type card struct {
		line int
		body strings.Builder
	}
	cards := map[string]*card{}
	var (
		cur   *card
		fence bool
	)
	sc := bufio.NewScanner(fh)
	// A table-only row carries its whole evidence; a few run past 64 KB.
	sc.Buffer(make([]byte, 0, 1<<16), 4<<20)
	for n := 1; sc.Scan(); n++ {
		ln := sc.Text()
		if fenceRe.MatchString(ln) {
			fence = !fence
		}
		if !fence {
			if m := cardRe.FindStringSubmatch(ln); m != nil {
				if cards[m[1]] != nil {
					l.dupes = append(l.dupes, fmt.Sprintf("%s at line %d", m[1], n))
					cur = &card{line: n}
					continue
				}
				cur = &card{line: n}
				cards[m[1]] = cur
				l.order = append(l.order, m[1])
				continue
			}
			if headingRe.MatchString(ln) {
				cur = nil
				continue
			}
		}
		if cur != nil {
			// Every other reader of claims/*.md takes an id-keyed table row for a
			// claim definition, so a card must not hold one.
			if !fence && strings.HasPrefix(strings.TrimSpace(ln), "|") {
				if cells := splitRow(ln); len(cells) > 0 {
					if id := strings.Trim(cells[0], "`* "); id != "" && idRe.FindString(id) == id {
						l.inner = append(l.inner, fmt.Sprintf("%s at line %d", id, n))
					}
				}
			}
			cur.body.WriteString(ln)
			cur.body.WriteByte('\n')
			continue
		}
		if fence || !strings.HasPrefix(strings.TrimSpace(ln), "|") {
			continue
		}
		cells := splitRow(ln)
		if len(cells) == 0 {
			continue
		}
		m := idRe.FindStringSubmatch(cells[0])
		if m == nil {
			continue
		}
		l.rows = append(l.rows, row{id: m[1], file: l.name, line: n, cells: cells, raw: ln})
	}
	if err := sc.Err(); err != nil {
		return l, fmt.Errorf("%s: %w", path, err)
	}
	seen := map[string]bool{}
	for i := range l.rows {
		r := &l.rows[i]
		seen[r.id] = true
		if c := cards[r.id]; c != nil {
			r.hasCard = true
			r.card = strings.Trim(c.body.String(), "\n")
			r.cardLine = c.line
		}
	}
	for _, id := range l.order {
		if !seen[id] {
			l.orphans = append(l.orphans, fmt.Sprintf("%s at line %d", id, cards[id].line))
		}
	}
	return l, nil
}

// splitRow splits a markdown table line on unescaped pipes that are not inside a
// backtick span. Claim text is full of `disp:a9c4 | 32 hits` and of escaped
// pipes, and a naive strings.Split puts the rest of the row in the wrong column.
func splitRow(ln string) []string {
	ln = strings.TrimSpace(ln)
	if !strings.HasPrefix(ln, "|") {
		return nil
	}
	var (
		cells []string
		cur   strings.Builder
		tick  bool
	)
	rs := []rune(ln)
	for i := 0; i < len(rs); i++ {
		switch {
		// Only `\|` and `\\` are escapes in a table cell. A lone backslash is data:
		// claim text carries Windows paths (`terrain\tile*.bmp`).
		case rs[i] == '\\' && i+1 < len(rs) && (rs[i+1] == '|' || rs[i+1] == '\\'):
			cur.WriteRune(rs[i+1])
			i++
		case rs[i] == '`':
			tick = !tick
			cur.WriteRune(rs[i])
		case rs[i] == '|' && !tick:
			cells = append(cells, strings.TrimSpace(cur.String()))
			cur.Reset()
		default:
			cur.WriteRune(rs[i])
		}
	}
	cells = append(cells, strings.TrimSpace(cur.String()))
	if len(cells) > 0 && cells[0] == "" {
		cells = cells[1:]
	}
	if len(cells) > 0 && cells[len(cells)-1] == "" {
		cells = cells[:len(cells)-1]
	}
	return cells
}

func printIDs(rows []row, ids []string) int {
	byID := map[string][]row{}
	for _, r := range rows {
		byID[r.id] = append(byID[r.id], r)
	}
	missing := 0
	for i, want := range ids {
		want = strings.ToUpper(strings.Trim(want, "`"))
		hits := byID[want]
		if len(hits) == 0 {
			fmt.Fprintf(os.Stderr, "claim: %s is not in any ledger\n", want)
			missing++
			continue
		}
		if i > 0 {
			fmt.Println()
		}
		// The claim first, then the retraction entries that qualify it.
		sort.SliceStable(hits, func(a, b int) bool {
			return hits[a].file != "retracted.md" && hits[b].file == "retracted.md"
		})
		for j, r := range hits {
			if j > 0 {
				fmt.Println()
			}
			printRow(r)
		}
	}
	if missing > 0 {
		return 1
	}
	return 0
}

// printRow prints one claim. The retraction state is the reason to cite a claim
// id rather than an experiment, so it is never something the caller asks for.
func printRow(r row) {
	mark := ""
	if r.file == "retracted.md" {
		mark = "  ** retraction entry: it corrects or narrows the claim above **"
	}
	fmt.Printf("%s  (claims/%s:%d)%s\n", r.id, r.file, r.line, mark)
	if r.hasCard && len(r.cells) == 5 {
		fmt.Printf("  %s\n", wrap(r.cells[1], 92, "  "))
		fmt.Printf("  Confidence: %s · Status: %s · Evidence: %s\n", r.cells[2], r.cells[3], plainLinks(r.cells[4]))
		if r.card != "" {
			fmt.Println()
			for _, ln := range strings.Split(r.card, "\n") {
				if ln == "" {
					fmt.Println()
					continue
				}
				fmt.Printf("  %s\n", ln)
			}
		}
		return
	}
	for i, c := range r.cells {
		// The id cell repeats the heading unless it carries a qualifier:
		// retracted.md scopes a row to "the writer-set clause only".
		if i == 0 && strings.Trim(c, "`* ") == r.id {
			continue
		}
		fmt.Printf("  %s\n", wrap(c, 92, "  "))
	}
}

func plainLinks(s string) string {
	return linkRe.ReplaceAllString(s, "$1")
}

func printLedger(rows []row, name string, full bool) int {
	name = strings.TrimSuffix(filepath.Base(name), ".md") + ".md"
	var hits []row
	for _, r := range rows {
		if r.file == name {
			hits = append(hits, r)
		}
	}
	if len(hits) == 0 {
		fmt.Fprintf(os.Stderr, "claim: no ledger %s\n", name)
		return 1
	}
	fmt.Printf("claims/%s — %d rows\n\n", name, len(hits))
	for _, r := range hits {
		if full {
			printRow(r)
			fmt.Println()
			continue
		}
		fmt.Printf("%-22s %-10s %s\n", r.id, status(r), summary(r, 96))
	}
	return 0
}

func printSearch(rows []row, pattern string, full bool) int {
	re, err := regexp.Compile("(?i)" + pattern)
	if err != nil {
		fmt.Fprintln(os.Stderr, "claim:", err)
		return 2
	}
	n := 0
	for _, r := range rows {
		// A card wraps its lines, so a phrase may straddle a line break.
		text := r.raw + "\n" + strings.Join(strings.Fields(r.card), " ")
		if !re.MatchString(text) {
			continue
		}
		n++
		if full {
			printRow(r)
			fmt.Println()
			continue
		}
		fmt.Printf("%-22s %-14s %-10s %s\n", r.id, strings.TrimSuffix(r.file, ".md"), status(r), summary(r, 78))
	}
	if n == 0 {
		fmt.Fprintf(os.Stderr, "claim: no row matches %q\n", pattern)
		return 1
	}
	fmt.Fprintf(os.Stderr, "\n%d rows. Re-run with -full for the evidence, or name the ids.\n", n)
	return 0
}

func printStats(root string, rows []row) {
	type agg struct {
		rows  int
		bytes int64
	}
	per := map[string]*agg{}
	for _, r := range rows {
		a := per[r.file]
		if a == nil {
			a = &agg{}
			per[r.file] = a
		}
		a.rows++
	}
	names := make([]string, 0, len(per))
	for n := range per {
		names = append(names, n)
		if fi, err := os.Stat(filepath.Join(root, n)); err == nil {
			per[n].bytes = fi.Size()
		}
	}
	sort.Slice(names, func(i, j int) bool { return per[names[i]].bytes > per[names[j]].bytes })

	var tb int64
	var tr int
	fmt.Printf("%-18s %10s %7s %10s\n", "ledger", "bytes", "rows", "bytes/row")
	for _, n := range names {
		a := per[n]
		fmt.Printf("%-18s %10d %7d %10d\n", n, a.bytes, a.rows, a.bytes/int64(max(a.rows, 1)))
		tb += a.bytes
		tr += a.rows
	}
	fmt.Printf("%-18s %10d %7d\n", "TOTAL", tb, tr)
	fmt.Fprintf(os.Stderr, "\nOpening one ledger costs its whole byte count. One claim is the size above.\n")
}

// lint checks every ledger that holds at least one card against the format in
// claims/registry.md. A ledger with rows and no cards is still table-only; it is
// counted, not failed. registry.md and retracted.md are not claim ledgers.
func lint(ledgers []ledger) int {
	var (
		bad      []string
		cardLeds int
		legacy   []string
		claims   int
	)
	fail := func(l ledger, line int, format string, args ...any) {
		bad = append(bad, fmt.Sprintf("claims/%s:%d: %s", l.name, line, fmt.Sprintf(format, args...)))
	}
	// kinds maps a claim id to the Kind cells of its retracted.md entries, so a
	// card's Status can be held to the corrections recorded against it.
	kinds := map[string][]string{}
	for _, l := range ledgers {
		if l.name != "retracted.md" {
			continue
		}
		for _, r := range l.rows {
			if len(r.cells) > 1 {
				kinds[r.id] = append(kinds[r.id], strings.ToUpper(r.cells[len(r.cells)-1]))
			}
		}
	}
	for _, l := range ledgers {
		if l.name == "registry.md" || l.name == "retracted.md" {
			continue
		}
		if len(l.order) == 0 && len(l.dupes) == 0 {
			if len(l.rows) > 0 {
				legacy = append(legacy, l.name)
			}
			continue
		}
		cardLeds++
		for _, o := range l.orphans {
			bad = append(bad, fmt.Sprintf("claims/%s: card %s has no index row", l.name, o))
		}
		for _, d := range l.dupes {
			bad = append(bad, fmt.Sprintf("claims/%s: second card for %s", l.name, d))
		}
		for _, d := range l.inner {
			bad = append(bad, fmt.Sprintf("claims/%s: table row keyed by claim id %s inside a card; other readers take it for a claim", l.name, d))
		}
		var rowOrder []string
		for _, r := range l.rows {
			claims++
			rowOrder = append(rowOrder, r.id)
			if !r.hasCard {
				fail(l, r.line, "%s has no card", r.id)
				continue
			}
			if len(r.cells) != 5 {
				fail(l, r.line, "%s: index row has %d cells, want 5", r.id, len(r.cells))
				continue
			}
			if n := utf8.RuneCountInString(r.cells[1]); n > headlineMax {
				fail(l, r.line, "%s: headline is %d characters, limit %d", r.id, n, headlineMax)
			}
			retracted := strings.HasPrefix(r.cells[3], "✖ retracted")
			if !gradeRe.MatchString(r.cells[2]) && !(retracted && r.cells[2] == "—") {
				fail(l, r.line, "%s: confidence %q is not High, Medium, Low or Unknown joined by \" / \"", r.id, r.cells[2])
			} else if !gradeOrder(r.cells[2]) {
				fail(l, r.line, "%s: confidence %q must name each grade once, in the order High, Medium, Low, Unknown", r.id, r.cells[2])
			}
			m := statusRe.FindStringSubmatch(r.cells[3])
			if m == nil {
				fail(l, r.line, "%s: status %q is not a marker with optional qualifiers", r.id, r.cells[3])
			} else if m[3] != "" {
				needsAmended := false
				for _, q := range strings.Split(m[3], ",") {
					q = strings.TrimSpace(q)
					amends, ok := qualifiers[q]
					if !ok {
						fail(l, r.line, "%s: status qualifier %q is not in claims/registry.md", r.id, q)
					}
					needsAmended = needsAmended || amends
				}
				if needsAmended && !strings.Contains(r.card, "**Amended.**") {
					fail(l, r.cardLine, "%s: status %q needs an **Amended.** paragraph", r.id, r.cells[3])
				}
			}
			if m != nil && m[1] != "✖ retracted" && m[3] == "" && strings.Contains(r.card, "**Amended.**") {
				fail(l, r.line, "%s: the card has an **Amended.** paragraph, so the status needs a qualifier", r.id)
			}
			if m != nil && m[1] != "✖ retracted" {
				for _, want := range statusFromRetractions(kinds[r.id]) {
					if (want == "" && m[3] == "") || (want != "" && !strings.Contains(m[3], want)) {
						if want == "" {
							want = "a qualifier"
						}
						fail(l, r.line, "%s: retracted.md records a correction against it; status %q needs %s", r.id, r.cells[3], want)
					}
				}
			}
			if !evidenceRe.MatchString(r.cells[4]) {
				fail(l, r.line, "%s: evidence names no experiment", r.id)
			}
			if len(r.card) > cardMax {
				fail(l, r.cardLine, "%s: card is %d bytes, limit %d", r.id, len(r.card), cardMax)
			}
		}
		if len(l.orphans) == 0 && len(l.dupes) == 0 && len(rowOrder) == len(l.order) {
			for i := range rowOrder {
				if rowOrder[i] != l.order[i] {
					bad = append(bad, fmt.Sprintf("claims/%s: cards are not in index order (row %s, card %s)", l.name, rowOrder[i], l.order[i]))
					break
				}
			}
		}
	}
	for _, b := range bad {
		fmt.Println(b)
	}
	if len(bad) > 0 {
		fmt.Printf("claim -check: %d finding(s) over %d card ledger(s)\n", len(bad), cardLeds)
		return 1
	}
	fmt.Printf("claim -check: ok (%d card ledger(s), %d claims; %d table-only ledger(s))\n", cardLeds, claims, len(legacy))
	return 0
}

// statusFromRetractions names what a Status must carry for the retracted.md
// entries recorded against its claim: any entry needs a qualifier, a
// SUPERSEDED entry needs "superseded", and a REFUTED or RETRACTED entry needs
// "partially retracted". A whole-claim ✖ retracted needs none of these.
func statusFromRetractions(kinds []string) []string {
	if len(kinds) == 0 {
		return nil
	}
	want := []string{""}
	for _, k := range kinds {
		if strings.Contains(k, "SUPERSEDED") && !contains(want, "superseded") {
			want = append(want, "superseded")
		}
		if (strings.Contains(k, "REFUTED") || strings.Contains(k, "RETRACTED")) && !contains(want, "partially retracted") {
			want = append(want, "partially retracted")
		}
	}
	return want
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

// gradeOrder reports whether a Confidence cell names each grade at most once,
// strongest first.
func gradeOrder(cell string) bool {
	rank := map[string]int{"High": 0, "Medium": 1, "Low": 2, "Unknown": 3}
	last := -1
	for _, g := range strings.Split(cell, " / ") {
		n, ok := rank[g]
		if !ok {
			return true // not a grade list; the pattern check reports it
		}
		if n <= last {
			return false
		}
		last = n
	}
	return true
}

func status(r row) string {
	for _, c := range r.cells {
		switch {
		case strings.Contains(c, "✖"):
			return "retracted"
		case strings.Contains(c, "✔"):
			return "promoted"
		case strings.Contains(c, "●"):
			return "active"
		}
	}
	if r.file == "retracted.md" {
		return "retracted"
	}
	return ""
}

// summary is the claim cell, flattened and cut, so an index costs one line per
// claim.
func summary(r row, n int) string {
	s := ""
	if len(r.cells) > 1 {
		s = r.cells[1]
	}
	s = strings.NewReplacer("**", "", "`", "", "\t", " ").Replace(s)
	s = strings.Join(strings.Fields(s), " ")
	if r.hasCard {
		return s
	}
	if utf8.RuneCountInString(s) > n {
		s = string([]rune(s)[:n-1]) + "…"
	}
	return s
}

func wrap(s string, width int, indent string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	var b strings.Builder
	col := 0
	for i, w := range words {
		if col > 0 && col+1+len(w) > width {
			b.WriteString("\n" + indent)
			col = 0
		} else if i > 0 {
			b.WriteString(" ")
			col++
		}
		b.WriteString(w)
		col += len(w)
	}
	return b.String()
}
