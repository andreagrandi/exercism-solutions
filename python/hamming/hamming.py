def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError('Sequences must be of the same length')

    distance = 0

    for i, v in enumerate(strand_a):
        if v != strand_b[i]:
            distance += 1

    return distance
