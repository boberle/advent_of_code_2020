from pprint import pprint


#POINT_COUNT = 30
#IGNORE = {2, 4, 7, 8, 11, 15, 22}
#POINT_COUNT = 15
#IGNORE = {2, 4, 7, 8}
POINT_COUNT = 157
IGNORE = {3, 4, 10, 11, 15, 16, 22, 23, 25, 26, 30, 31, 36, 37, 43, 44, 46, 47, 50, 51, 56, 57, 61, 62, 67, 68, 72, 73, 79, 80, 82, 83, 88, 89, 95, 96, 101, 102, 107, 108, 114, 115, 121, 122, 126, 127, 133, 134, 136, 137, 142, 143, 145, 146, 152, 153, 155, 156}


def compute_paths(point_count):
    paths = [list() for i in range(point_count + 1)]
    paths[0].append([0])
    for i in range(1, point_count + 1):
        if i in IGNORE:
            continue
        for back in range(1, 4):
            if i - back >= 0:
                for path in paths[i-back]:
                    new_path = list(path) + [i]
                    paths[i].append(new_path)
    return paths


cache = dict()

def fib3(n):
    if n in IGNORE:
        return 0
    try:
        return cache[n]
    except KeyError:
        if n == 0:
            res = 1
        elif n == 1:
            res = 1
        elif n == 2:
            res = 2
        else:
            res = fib3(n-1) + fib3(n-2) + fib3(n-3)
        cache[n] = res
        return res



def main():
    fib3_res = fib3(POINT_COUNT)
    print(fib3_res)
    return

    paths = compute_paths(POINT_COUNT)
    #pprint(paths)
    for i, path_set in enumerate(paths):
        print(i, len(path_set))
    print(fib3_res == len(path_set))


if __name__ == "__main__":
    main()
