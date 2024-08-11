# AoC 2020, day 7

Solution to the test file:

```
Part 1: number of shiny gold ancestors: 4
Part 2: number of bags in the shiny gold bag: 32
```

and the second test file:

```
Part 1: number of shiny gold ancestors: 0
Part 2: number of bags in the shiny gold bag: 126
```

## How does it work?

For the first part, we go with a simple approach: we build the tree of all bags (a bag color is a structure containing children bag colors). Then we look for the ancestors.

Then, for the second, part, we just count descend the tree of the shiny gold bag, recursively.