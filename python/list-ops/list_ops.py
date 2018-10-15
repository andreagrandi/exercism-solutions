def append(xs, ys):
    for y in ys:
        xs.insert(len(xs), y)
    return xs


def concat(lists):
    concat_list = []
    for l in lists:
        concat_list += l
    return concat_list


def filter_clone(function, xs):
    filtered_list = []
    for x in xs:
        if function(x):
            filtered_list.insert(len(filtered_list), x)
    return filtered_list


def length(xs):
    list_length = 0
    for x in xs:
        list_length += 1
    return list_length


def map_clone(function, xs):
    mapped_list = []
    for x in xs:
        mapped_list.insert(len(mapped_list), function(x))
    return mapped_list


def foldl(function, xs, acc):
    for x in xs:
        acc = function(acc, x)
    return acc


def foldr(function, xs, acc):
    xs = xs[::-1]
    for x in xs:
        acc = function(x, acc)
    return acc


def reverse(xs):
    return xs[::-1]
