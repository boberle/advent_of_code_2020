# AoC 2020, day 19

Solution to the first test file (only for the first part):

```
Number of valid messages (part 1): 2
```

Solution to the second test file:

```
Number of valid messages (part 1): 3
Number of valid messages (part 2): 12
```

## How to solve?

The first part is pretty straightforward, you just need to build the rule tree, then check if each string matches the top rule (rule 0).

This works pretty well since it appears that there is no ambiguity in the datasets.

For the second part, as the statement reads, we only need to handle the simple rule that we have, not to write a formal grammar parser.

The rules are as follows:

```
0: 8 11
8: 42 | 42 8 (instead of 8: 42)
11: 42 31 | 42 11 31 (instead of 11: 42 11)
```

These rules don't change, whatever the dataset is provided.

A quick analysis shows that rule 8 is just a variable number of repetition of rule 42, which means:

```
8: 42 | 42 42 | 42 42 42 | ...
```

and rule 11 is also a variable number of repetition of rule 42 and 31:

```
11: 42 31 | 42 42 31 31 | 42 42 42 31 31 31 | ...
```

Knowing that, you really just need to loop over rules 42 and 42/31 a reasonable amount of times (enough to cover all possibilities according to the length of the string to match).

Because we don't know the number of repetition of rules 8 and 11, we just loop over 11 inside 8 in order to test all possible possibilities. This works well since there is no ambiguity.
