fn = {}

fn['fn1'] = lambda x, y, z: x + y + z
fn['fn2'] = lambda : "No arg need"

print(fn['fn1']('1', '23', '456'))
print(fn['fn2']())
