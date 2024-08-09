# AoC 2020, day 6

Solution to the test file:

```
Part 1: number of questions anyone yes-answered in all group: 11
Part 2: number of questions everyone yes-answered in all group: 6
```

## How does it work?

For the first part, just build the set of each question in each group, and sum up the lengths of those sets.

For the second part, for each group, we build a map `question -> number of time answered in the group`. Then we count the number of questions answered by all the members of the group, that is, the number of times the number of members of the group appears has a value in the map.