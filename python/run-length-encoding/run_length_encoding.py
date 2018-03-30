import re
def decode(string):
    result = ""
    for count, letter in re.findall(r'(\d*)([A-Za-z ])', string):
        if count == '':
            count = 1
        result += letter * int(count) 
    return result


def encode(string):
    result = ""
    for repeated, _ in re.findall(r'(([A-Za-z ])\2*)', string):
        l = len(repeated)
        if l > 1:
            result += str(l)
        result += repeated[0] 
    return result
        
