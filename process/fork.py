import os
from time import sleep

try:
    pid = os.fork()
except OSError as e:
    print(e)

if pid > 0:
    print("This is parent")
    print("Parent get the child's pid")
else:
    print("This is Child")
    print("Child get 0")
print(pid)
