from math import sqrt, pow

def solution(n):
    sqrt_ = int(sqrt(n))
    if sqrt_ * sqrt_ == n:
        return int(pow((sqrt_+1),2))
    else:
        return -1

print(solution(121))