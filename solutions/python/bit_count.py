def count_bit(n):
    ret = 0
    b = []
    while n > 0:
        d = n % 2
        b.append(d)
        n = n//2
    b.reverse()
    for i in b:
        if i == 1:
            ret += 1
    
    return ret

print(count_bit(1234))