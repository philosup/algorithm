N, M = map(int, input().split())
num = list(map(int, input().split()))

def DFS(idx:int, sum:int, depth:int):
    # print(f'{idx},{sum},{depth}')
    if depth == 3:
        return sum
    max_sum = 0
    for i in range(idx, len(num)):
        _sum = DFS(i+1, sum + num[i], depth + 1)
        if _sum <= M and _sum > max_sum:
            max_sum = _sum
    return max_sum

print(DFS(0,0,0))