class Node:
    def __init__(self, item, left, right):
        self.item = item
        self.left = left
        self.right = right

def printif(node, bPrint):
    if bPrint: print(node.item, end='')

out = []

def traverse(node, mode):
    old_len = len(out)
    left = 0
    if node.left != '.':
        left = traverse(tree[node.left], mode)
    if node.right != '.':
        traverse(tree[node.right], mode)

    if old_len > left: left = old_len
    if mode == -1:  out.insert(old_len, node.item)
    elif mode == 0: out.insert(left, node.item)       
    else:           out.append(node.item)

    return len(out)

def print_out(bClear = True):
    print(*out, sep='')
    if bClear: out.clear()

n = int(input())
tree = {}  # dictionary
for _ in range(n):
    # 입력값 가져와서 key, value 로 매핑
    node, l, r = map(str, input().split())
    tree[node] = Node(item=node, left=l, right=r)


traverse(tree['A'], -1)
print_out()
traverse(tree['A'], 0)
print_out()
traverse(tree['A'], 1)
print_out()
