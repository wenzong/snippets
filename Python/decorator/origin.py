def do2048(func):
    def inner_func(*args, **kwargs):
        for i in xrange(2048):
            func(*args, **kwargs)
    return inner_func

def say_hello():
    print("hello")

new_func = do2048(say_hello)

new_func() # do2048(say_hello)()
