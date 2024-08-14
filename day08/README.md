# AoC 2020, day 8

Solution to the test file:

```
Part 1: the accumulator value is 5
Part 2: the accumulator value is 8
```

## How does it work?

For the first part, we read all the instruction and execute them. When an instruction is visited for the second time, it is an indication that we will run an infinite loop, so we stop.

For the second part, we use a brute force approach: for each `jmp` or `nop` instruction, we change it to `nop` or `jmp`, respectively, and try to run the program. If it finishes, then we have find the corrupted instruction.