import functools
fn = {}

def decorator_func(name):
    def wrapper(func):
        fn[name] = func
    return wrapper

@decorator_func("fn1")
def fn1(arg1, arg2, arg3):
    return arg1 + arg2 + arg3

@decorator_func("fn2")
def fn2():
    return "No arg need"

print(fn['fn1']('1', '23', '456'))
print(fn['fn2']())

print(fn1)  # None
print(fn2)  # None
