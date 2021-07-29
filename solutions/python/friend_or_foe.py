def friend(x):
    frnds = list()
    for i in x:
        if len(i) == 4:
            frnds.append(i)

    return frnds

print(friend(["Ryan", "Kieran", "Mark",]))