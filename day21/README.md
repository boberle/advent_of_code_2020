# AoC 2020, day 19

Solution to the test file:

```
How many times non allergenic ingredient appears (part 1): 5
Canonical list (part 2): mxmxvkd,sqjhc,fvjkl
```

## How to solve?

We need to make a mapping between ingredient names (eg `mxmxvkd`) and the allergen (eg `dairy`). To do that, we pair allergen with ingredient list, like this:

```
dairy: appears in "mxmxvkd kfcds sqjhc nhms" and in "trh fvjkl sbzzf mxmxvkd"
fist: appears in "mxmxvkd kfcds sqjhc nhms" and "sqjhc mxmxvkd sbzzf"
soy: appears in "sqjhc fvjkl"
```

and then we find the common ingredient that appears in all the list for a given allergen:

```
"mxmxvkd" appears in all list for "dairy"
"sqjhc" appears in all list for "fish"
```

and then we can deduce that "soy" is "fvjkl".
