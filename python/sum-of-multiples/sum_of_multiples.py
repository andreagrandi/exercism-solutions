def sum_of_multiples(limit, multiples):
    sum = 0
    multiples_of_limit = []

    for i in range(1, limit):
        for multiple in multiples:
            if i % multiple == 0:
                multiples_of_limit.append(i)

    multiples_of_limit = set(multiples_of_limit)

    for multiple in multiples_of_limit:
        sum += multiple

    return sum
