def verify(isbn, length=10):
    isbn = isbn.replace('-', '')
    if len(isbn) != length:
        return False
    check = 0
    for i, char in enumerate(isbn):
        if  i == length-1 and char == 'X':
            digit = 10
        else:
            digit = ord(char) - ord('0')

        if digit < 0 or digit > 10:
            return False

        check += (length - i) * digit
    return check % (length + 1) == 0
        

