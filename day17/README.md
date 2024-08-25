# AoC 2020, day 17

Solution to the test file:

```
active cubes after 6 cycles (part 1): 112
active hypercubes after 6 cycles (part 2): 848
```

## How does it work?

Both parts are solved in the same manners, for the second part, just add a dimension (x, y, z, w instead of just x, y, z).

We keep a list of active cubes (this is the input) L. For each cycle:

- prepare a list of new active cubes (the cubes that will be active at the end of the cycle)
- prepare a map I of inactive cubes (keys are cube positions and values are the number of cubes that have that inactive cube as neighbor)
- for each cube C in L:
  - for each neighbor N, if N is:
    - in L, then it is active and add one to the counter C
    - not in L, then it is inactive, and add it to I (the key is N, and add one to the value)
  - if the count of active neighbors is 2 or 3, add C to the list of new actives cubes
- now for the inactive cubes that will go active: walk through I and if the value is 3, then add the key to the new active cubes.