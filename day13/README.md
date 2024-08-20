# AoC 2020, day 13

Solution to the test file:

```
Earliest bus * waiting time (part 1): 295
Earliest timestamp (part 2): 1068781
```

Here is the solution in python, along side the Go solution in the corresponding files:

```python
def parse_data(data):
    items = []
    for i, dat in enumerate(data.split(",")):
        dat = dat.strip()
        if dat.isdigit():
            a = int(dat)
            items.append(((a - i) % a, a))
    return items

def part1(ids, t):
    min_waiting_time = None
    min_id = None
    for id in ids:
        waiting_time = (t // id + 1) * id - t
        if min_waiting_time is None or waiting_time < min_waiting_time:
            min_waiting_time = waiting_time
            min_id = id
    return min_waiting_time * min_id

def inv_mod(a, m):
    for i in range(0, m):
        if (a * i) % m == 1:
            return i
    return None

def solve(*items):
    M = 1
    for item in items:
        M *= item[1]

    x0 = 0
    for a, m in items:
        Mi = M // m
        x0 += a * Mi * inv_mod(Mi, m)

    x = x0 % M
    return x


items = parse_data("7,13,x,x,59,x,31,19")
print("part 1:", part1(map(lambda x: x[1], items), 939))
print("part 2:", solve(*items))
```

## How it works?

For the first part, we want to minimize the waiting time. So we compute for each bus the waiting time:

```python
t = 939
for bus_id in bus_ids:
    next_departure = (t // id + 1) * id
    waiting_time = next_departure - t
```

For the second part, using another example given in the statement: `17,x,13,19`, we just need to recognize that what we want is:

```
| (x + 0) % 17 = 0 -> x % 17 = (17 - 0) % 17 -> x % 17 = 0
| (x + 2) % 13 = 0 -> x % 13 = (13 - 2) % 13 -> x % 13 = 11
| (x + 3) % 19 = 0 -> x % 19 = (19 - 3) % 19 -> x % 19 = 16
```

We thus have the set congruences:

```
| x ≡ 0 (mod 17)
| x ≡ 11 (mod 13)
| x ≡ 16 (mod 19)
```

And we can find x with the Chinese remainder theorem (note that for each, `GCD(a, m) = 1`).  The code is really just a very very naive implementation of the theorem.

You could also brute force the thing (using the first example in the statement):

```c
#include <stdlib.h>
#include <stdio.h>

void dataset_test();

int main() {
   dataset_test();
   return 0;
}

void dataset_test() {
   long long unsigned t = 0, c = 0;
   while (1) {
      t += 59;
      if (
         (t-4) % 7 == 0 &&
         (t-3) % 13 == 0 &&
         t % 59 == 0 &&
         (t+2) % 31 == 0 &&
         (t+3) % 19 == 0
      ) {
         printf("found: %llu\n", t - 4);
         break;
      }
   }
}
```

The result is `1068781`. This brute force approach works even for the real datasets, but you have to wait a few hours.