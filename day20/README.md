# AoC 2020, day 20

Solution to the test file:

```
Ids of the corners multiplied together (part 1): 20899048083289
Found 2 monsters, so there are 303 - 30 = 273 points in the image that are not sea monsters (part 2)
```

How it works?

For the first part, there is a quick solution in the file `run.go.part1`:

We register in a set all the sides that match another one (exactly, or in reverse). Then we walk through all the tiles, and count tiles that have exactly 2 matching sides (that is, tiles that have 2 sides in the set). These are the corners.

For the second part, we need to assemble the whole image. This is done with the following algorithm:

1. take randomly a tile
2. for each side, look at all the other tiles: rotate/flip each of them to check if it is a matching tile for the corresponding side. For example, look for the right tile (if any) by trying to rotate/flip all other tile. If found, then pair the tiles together.
3. once you have found a tile to pair with one side, repeat step 2 for that tile.

For example:

- take tile A
- look for the tile at the right of A (if any): try to rotate/flip tile B, tile C, etc. to see if it goes to the right of B
- let's say that tile D goes to the right of A
- look for the tile at the right of D. If there is, then look for the tile at the right of the new tile, etc.
- then look for the tile at the bottom of A, and repeat
- same for the tile at the left and top of A

Once you have build the image, you can just look for the monsters in each possible rotation/flipping of the entire image.