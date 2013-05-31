import sys

def fib(n):
    if n<2:
        return 1
    else:
        return fib(n-1) + fib(n-2)

n = sys.argv[1] if len(sys.argv) >= 2 else 13
print 'fib(%d)=%d' % (int(n), fib(int(n)))

