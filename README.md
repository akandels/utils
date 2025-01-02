A collection of small utilities in a single binary for common command line use.

## split-group

Used to split log lines into groups which are then counted and compared to each
other to identify patterns or anomalies; used with STDIN or a file piped in as
input.

Example usage:

`echo "red dog\nblue dog\nyellow dog\npurple dog\n" | ./utils split-group "(red|blue)" yellow`

Results:

```
Matches:
NO_MATCH:   0 (0%)
(red|blue): 2 (66%)
yellow:     1 (33%)
```

Flags:

| Flag | Description | Default |
| ---- | ----------- | ------- |
| --log-non-matches, -x | Log non-matching lines, helpful for gradually adding patterns initially | Off |
| --multi-match, -m | Allow lines to match multiple patterns | Off |

## split-match

Given a regular expression with a single capture group, processes log lines from STDIN
and groups by the value captured. Useful for counting log lines by IP address, customer, etc.

Example usage:

`echo '{"company_id": 123}\n{"company_id: 124}\n{"company_id": 123}\n' | ./utils split-match 'company_id":\s*(\d+)'`

Results:

```
Matches:
NO_MATCH: 0
123:      2 (66%)
124:      1 (33%)
```
