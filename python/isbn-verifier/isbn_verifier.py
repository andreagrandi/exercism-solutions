def verify(isbn):
    # Remove all '-' characters
    isbn = isbn.replace('-', '')

    # Check for valid length of the code
    if len(isbn) != 10:
        return False

    # Check for valid check caracter at the end
    if isbn[-1] not in ['8', 'X']:
        return False

    valid_characters = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'X']

    # Convert the string to an array of numbers (converting the check too)
    digits = []
    for c in isbn:
        if c not in valid_characters:
            return False

        if c == 'X':
            digits.append(10)
        else:
            digits.append(int(c))

    # Calculate the validation using the formula
    # (x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) mod 11 == 0
    validation = (
        digits[0] * 10 + digits[1] * 9 + digits[2] * 8 + digits[3] * 7 + digits[4] * 6 + digits[5] * 5 +
        digits[6] * 4 + digits[7] * 3 + digits[8] * 2 + digits[9] * 1) % 11

    return validation == 0
