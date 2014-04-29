import functools
fn = {}

def decorator_func(func):
    fn[func.__name__] = func

@decorator_func
def fn1(arg1, arg2, arg3):
    return arg1 + arg2 + arg3

@decorator_func
def fn2():
    return "No arg need"

print(fn['fn1']('1', '23', '456'))
print(fn['fn2']())

print(fn1)
print(fn2)
