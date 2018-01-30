
"""

Source: https://talks.golang.org/2013/go4python.slide#16
"""

def fib(n):
    a, b = 0, 1
    for i in range(n):
        a, b = b, a + b
        yield a

for x in fib(10):
    print(x)
