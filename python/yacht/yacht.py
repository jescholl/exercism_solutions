# Score categories
# Change the values as you see fit
YACHT = 11111
ONES = 1
TWOS = 2
THREES = 3
FOURS = 4
FIVES = 5
SIXES = 6
FULL_HOUSE = 11222
FOUR_OF_A_KIND = 1111
LITTLE_STRAIGHT = 1234
BIG_STRAIGHT = 12345
CHOICE = 11233


def score(dice, category):
    if category == YACHT:
        return yacht(dice)
    elif category == FULL_HOUSE:
        return full_house(dice)
    elif category == FOUR_OF_A_KIND:
        return four_of_a_kind(dice)
    elif category == LITTLE_STRAIGHT:
        return straight(dice, 1)
    elif category == BIG_STRAIGHT:
        return straight(dice, 2)
    elif category == CHOICE:
        return choice(dice)
    else:
        return nums(dice, category)

def nums(dice, num):
    score = 0
    for die in dice:
        if die == num:
            score += die
    return score

def yacht(dice):
    for i in range(1, 5):
        if dice[i] != dice[0]:
            return 0
    return 50

def full_house(dice):
    d = {x:dice.count(x) for x in dice}
    if 3 in d.values() and 2 in d.values():
        return sum(dice)
    return 0

def four_of_a_kind(dice):
    d = {x:dice.count(x) for x in dice}
    for k,v in d.items():
        if v >= 4:
            return k*4
    return 0

def straight(dice, starts_at):
    dice.sort()
    print(dice, "!=", list(range(starts_at, starts_at+5)))
    if dice == list(range(starts_at, starts_at+5)):
        return 30
    return 0

def choice(dice):
    return sum(dice)
