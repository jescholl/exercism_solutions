from string import ascii_lowercase

def is_pangram(sentence):
    letters = set(sentence.lower())
    return len(set(ascii_lowercase) - letters) == 0
