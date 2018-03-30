def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        print("raising excpetion")
        raise ValueError("Inputs must be of equal length")
    distance = 0
    for idx, char in enumerate(strand_a):
        if char != strand_b[idx]:
            distance += 1
    return distance
