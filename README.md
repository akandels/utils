A collection of small utilities in a single binary for common command line use.

## split-group

Used to split log lines into groups which are then counted and compared to each
other to identify patterns or anomalies; used with STDIN or a file piped in as
input.

Example, given the following log lines:

- A red dog walked into the park.
- A blue dog left the park.
- A yellow dog chased a purple dog.

`split-group "(red|blue)" yellow`

> Matches:
> NO_MATCH: 0
> (red|blue): 2 (66%)
> yellow: 1 (33%)

Flags:

| Flag | Description | Default |
| ---- | ----------- | ------- |
| --log-non-matches, -x | Log non-matching lines, helpful for gradually adding patterns initially | Off |
| --multi-match, -m | Allow lines to match multiple patterns | Off |