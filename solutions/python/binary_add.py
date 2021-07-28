def binary_add(a,b):
    res = 0
    z = a + b
    b_ = []
    while z > 0:
        d = z % 2
        b_.append(d)
        z = z//2
    b_.reverse()

    res = ''.join(str(i) for i in b_)
    
    return res

print(binary_add(5, 9))