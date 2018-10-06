def is_isogram(string):
    # Removes spaces and hyphens
    string = string.replace(' ', '')
    string = string.replace('-', '')

    # Transform the string to lower case
    string = string.lower()

    return len(list(string)) == len(set(string))
