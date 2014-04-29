fn = {}

def fn1(arg1, arg2, arg3):
    return arg1 + arg2 + arg3

def fn2():
    return "No arg need"

fn['fn1'] = fn1
fn['fn2'] = fn2

print(fn['fn1']('1', '23', '456'))
print(fn['fn2']())
