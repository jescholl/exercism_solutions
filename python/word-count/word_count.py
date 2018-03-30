from collections import defaultdict
import re

def word_count(phrase):
    words = re.findall(r"[a-z0-9']+", phrase.lower())
    count = defaultdict(int)
    for word in words:
        word = word.strip('"\'')
        count[word] += 1
    return count
