def is_text_uppercase(text):
    is_uppercase = False

    for c in text:
        if c.isalpha():
            if c == c.upper():
                is_uppercase = True
            else:
                return False
    return is_uppercase


def hey(phrase):
    # Remove white spaces at the beginning and at the end
    phrase = phrase.strip()

    # Check if there are any characters
    number_of_chars = 0

    for char in phrase:
        if char.isalnum():
            number_of_chars += 1

    if number_of_chars == 0:
        if len(phrase) == 0:
            return "Fine. Be that way!"

    if phrase[-1] == '?' and is_text_uppercase(phrase):
        return "Calm down, I know what I'm doing!"

    if phrase[-1] == '?' and not is_text_uppercase(phrase):
        return "Sure."

    if is_text_uppercase(phrase) and number_of_chars != 0:
        return "Whoa, chill out!"

    return "Whatever."
