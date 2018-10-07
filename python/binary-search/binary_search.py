def binary_search(list_of_numbers, number):
    if len(list_of_numbers) == 0:
        raise ValueError('empty array')

    start_index = 0
    stop_index = len(list_of_numbers) - 1

    while(start_index <= stop_index):
        mid_index = (start_index + stop_index) // 2

        if list_of_numbers[mid_index] == number:
            return mid_index

        if number > list_of_numbers[mid_index]:
            start_index = mid_index + 1
        else:
            stop_index = mid_index - 1

    raise ValueError('element not in the list')
