def is_isogram(string):
    allowed_duplicates = str.maketrans(dict.fromkeys(' -'))
    string = string.lower().translate(allowed_duplicates)
    s = set(string)
    if len(s) == len(string):
        return True
    else:
        return False

