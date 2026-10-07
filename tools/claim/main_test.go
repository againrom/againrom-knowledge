package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadThreeAndFourDigitClaimIDs(t *testing.T) {
	dir := t.TempDir()
	const ledger = "| ID | Claim |\n|---|---|\n| TOWN-186 | old claim |\n| SAV-CASTCONT-1006 | new claim |\n| SPR16A-FONT-018 | numbered topic |\n"
	if err := os.WriteFile(filepath.Join(dir, "sample.md"), []byte(ledger), 0600); err != nil {
		t.Fatal(err)
	}
	rows, err := load(dir)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"TOWN-186", "SAV-CASTCONT-1006", "SPR16A-FONT-018"}
	if len(rows) != len(want) {
		t.Fatalf("got %d claims, want %d", len(rows), len(want))
	}
	for i := range want {
		if rows[i].id != want[i] {
			t.Fatalf("row %d = %q, want %q", i, rows[i].id, want[i])
		}
	}
}

const cardLedger = `# X — sample

## Topic

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| X-A-001 | First headline. | High | ● active | [EXP-0001](../experiments/EXP-0001-a/) |
| X-A-002 | Second headline. | High / Medium | ● active (amended) | [EXP-0002](../experiments/EXP-0002-b/) |

### X-A-001

Body line.

| value | a table inside a card | is | not | a row |

**Confidence.** High.

### X-A-002

**Amended.** A clause changed.

## Next topic
`

func TestCardsAttachToRows(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "x.md"), []byte(cardLedger), 0600); err != nil {
		t.Fatal(err)
	}
	ledgers, err := loadLedgers(dir)
	if err != nil {
		t.Fatal(err)
	}
	rows := ledgers[0].rows
	if len(rows) != 2 {
		t.Fatalf("got %d rows, want 2 (a table inside a card is not a row)", len(rows))
	}
	if !rows[0].hasCard || rows[0].cardLine != 10 {
		t.Fatalf("X-A-001 card = %v at line %d, want a card at line 10", rows[0].hasCard, rows[0].cardLine)
	}
	if want := "Body line."; rows[0].card[:len(want)] != want {
		t.Fatalf("X-A-001 card starts %q", rows[0].card)
	}
	if got := rows[1].card; got != "**Amended.** A clause changed." {
		t.Fatalf("X-A-002 card = %q; a heading must end the card", got)
	}
	if code := lint(ledgers); code != 0 {
		t.Fatalf("lint = %d on a valid card ledger", code)
	}
}

func TestLintRefusesMalformedCards(t *testing.T) {
	cases := map[string]string{
		"missing card":   "### X-A-002\n\n**Amended.** A clause changed.\n",
		"orphan card":    "## Next topic\n",
		"bad grade":      "| High / Medium |",
		"bad qualifier":  "(amended)",
		"no amended":     "**Amended.** A clause changed.",
		"no evidence":    "[EXP-0001](../experiments/EXP-0001-a/)",
		"long headline":  "First headline.",
		"order":          "### X-A-001",
		"id row in card": "| value | a table inside a card |",
		"repeated grade": "| High / Medium |",
	}
	replace := map[string]string{
		"missing card":   "",
		"orphan card":    "### X-A-003\n\nstray\n\n## Next topic\n",
		"bad grade":      "| Mostly high |",
		"bad qualifier":  "(corrected)",
		"no amended":     "A clause changed.",
		"no evidence":    "none",
		"long headline":  strings.Repeat("x", headlineMax+1),
		"order":          "### X-A-000",
		"id row in card": "| X-A-001 | a table inside a card |",
		"repeated grade": "| High / High / Medium |",
	}
	for name, old := range cases {
		text := strings.Replace(cardLedger, old, replace[name], 1)
		if name == "order" {
			// Swap the two cards so they no longer follow the index.
			a := strings.Index(cardLedger, "### X-A-001")
			b := strings.Index(cardLedger, "### X-A-002")
			c := strings.Index(cardLedger, "## Next topic")
			text = cardLedger[:a] + cardLedger[b:c] + "\n" + cardLedger[a:b] + cardLedger[c:]
		}
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "x.md"), []byte(text), 0600); err != nil {
			t.Fatal(err)
		}
		ledgers, err := loadLedgers(dir)
		if err != nil {
			t.Fatal(err)
		}
		if code := lint(ledgers); code == 0 {
			t.Errorf("%s: lint passed", name)
		}
	}
}
