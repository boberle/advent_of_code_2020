# AoC 2020, day 25

Solution to the test file:

```
14897079
```

Well, nothing too complicated. Here is the solution in Python:

```python
card_pk = 5764801
door_pk = 17807724

def transform(subject_number, *, public_key=None, loop_size=None):
    assert public_key or loop_size
    size = 0
    value = 1
    while size := size + 1:
        value = (value * subject_number) % 20201227
        if value == public_key or size == loop_size:
            return size, value

card_loop_size, _ = transform(7, public_key=card_pk)
_, encryption_key = transform(door_pk, loop_size=card_loop_size)
print(encryption_key)
```

You just need to follow the instructions and do some brute forcing to perform the transformation.