# Contributing

For a factual correction, identify the claim ID, the disputed clause and the version
being read. State whether the basis is your own observation, general platform knowledge
or an external source. Do not paste another game's implementation, decompiled function,
resource file or an unreviewed table into an issue or pull request.

Game-specific claims follow the upstream research process and source policy. This
repository is not a second independently edited factual authority. Publication wording,
reader fixes and review tooling can be proposed here; their accepted transformation must
be retained by the exporter so the next snapshot does not undo the change.

Do not increase confidence or remove an Unknown while rewriting for publication. Keep
permanent IDs, evidence identifiers, and the meaning and scope of retractions. Exact
field widths, constants and behavioural qualifications are not expendable prose.

New tooling and fixtures must be independently authored. Test data should be synthetic
and clearly identified as such, not extracted game content disguised as a constant.
Check [THIRD_PARTY.md](THIRD_PARTY.md) and [PUBLICATION.md](PUBLICATION.md) before adding
any third-party material. Do not claim rights you do not have.

Before proposing changes, run the synthetic tests and inspect the publication report:

```sh
python3 -m unittest discover -s scripts -p 'test_*.py'
python3 scripts/publication_check.py --json
```

An issue or pull request is public on a public repository. Do not include receipts,
private credentials, personal logs, original assets or confidential source records.
