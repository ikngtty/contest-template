n = int(input())
x, y, z = (int(word) for word in input().split())
a = []
b = []
for i in range(n):
    ai, bi = (int(word) for word in input().split())
    a.append(ai)
    b.append(bi)

total = sum(a) + sum(b) + x + y + z

print(total)
