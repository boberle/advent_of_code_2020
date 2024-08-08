# AoC 2020, day 5

Solution to the test file:

```
Part 1: Highest seat id is 820 among 4 numbers
```

There is no test data for the part 2.


## How does it work?

For the first part, we just need to sum up the powers of 2 that are in the second half.

For example, to find the row, we have `F` (front) and `B` (back). We just sum up powers of 2 where there is a `B`, but in reverse order. So for `FBFBBFF`:

- we reverse the string: `FFBBFBF`,
- then find the positions of `B`, starting at 0: 2, 3, 5,
- then sum up the power of 2: `2^1 + 2^3 + 2^4 = 44 `

For the second parts, we look for the only seat that has an ID surrounded with ID-1 and ID+1.