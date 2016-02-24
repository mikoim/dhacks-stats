# dhacks-stats
Print repository statistics in 'dhacks' organization to stdout.
Maybe it's useful for managers.

## Usage

```
$ git clone git@github.com:mikoim/dhacks-stats.git
$ cd dhacks-stats
$ gb build && ./bin/ds
dhacks/team-a		0
dhacks/team-b	Java	5996
dhacks/team-c		0
dhacks/team-d	JavaScript	7
dhacks/team-e		0
dhacks/team-f	Ruby	45
dhacks/team-g		0
dhacks/team-h	Python	32
dhacks/team-i		0
dhacks/team-j		0
```

## Output Format

```
<Repository Name>\t<Language>\t<Size (KB)>\n
```
