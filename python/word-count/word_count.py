def word_count(phrase):
    # Replace non alpha numeric chars with spaces
    words_alpha_numeric = ""

    for c in phrase:
        if c == '\'':
            words_alpha_numeric += '\''
        elif c.isalnum():
            words_alpha_numeric += c.lower()
        else:
            words_alpha_numeric += ' '

    words = words_alpha_numeric.split()
    word_count = {}

    for word in words:
        # Check if it's a quoted word first
        if word[0] == '\'' and word[-1] == '\'':
            word = word.replace('\'', '')

        if word not in word_count:
            word_count[word] = 1
        else:
            word_count[word] += 1

    return word_count
