import sys

n = int(sys.stdin.readline())
tree = {}
for i in range(1,n+1):
    tree[i] = []

for i in range(n-1):
    v1, v2 = map(int, sys.stdin.readline().split())
    tree[v1].append(v2)
    tree[v2].append(v1)

q = [1]
ans = [0]*(n+1)
ans[1] = 1

while q:
    parent = q.pop(0)
    for i in tree[parent]:
        if not ans[i]:
            ans[i] = parent
            q.append(i)

print(*ans[2:], sep='\n')
# for i in range(2, n+1):
#     print(ans[i])