import re
def hey(phrase):
    phrase = phrase.strip()
    if phrase.isupper():
        if phrase.endswith("?"):
            return "Calm down, I know what I'm doing!"
        return "Whoa, chill out!"
    if phrase.endswith("?"):
        return "Sure."
    if not phrase:
        return "Fine. Be that way!"
    return "Whatever."
