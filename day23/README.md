# AoC 2020, day 23

Solution to the test file:

```
The order of the cups after cup 1 and 100 moves is (part 1): 67384529
The multiplied labels of the 2 cups after cup 1 are (part 2): 149245887792
```

How it works?

The solution I came up with is to use:

- a linked list for the cups,
- an array of pointers to the cups for the labels: it is mapping labels -> cups, the labels being the indices (the index 0 is ignored),
- a variable that holds a pointer to the current cup.

The linked list just reference the cup id and a pointer to the next cup. The last cup points to the first cup.

At each turn:

- we extract the 3 cups following the current cup by changing the pointers in the linked list,
- we look for the destination cup in the array (it's just the index below the current cup label),
- we add the 3 cups back in the link list by changing the pointers.

The final calculations are done by looking for the cup 1, which is at index 1 in the array.

