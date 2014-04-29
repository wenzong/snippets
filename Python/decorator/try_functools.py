from functools import wraps

def decorator1(func):
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)
    return wrapper

@decorator1
def func1(*args, **kwargs):
    """Doc string"""
    return

print(func1)              # <function wrapper at xxxxxx>
print(func1.__name__)     # wrapper
print(func1.__doc__)      # None


def decorator2(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)
    return wrapper

@decorator2
def func2(*args, **kwargs):
    """Doc string"""
    return

print(func2)              # <function func2 at xxxxxx>
print(func2.__name__)     # func2
print(func2.__doc__)      # Doc string
