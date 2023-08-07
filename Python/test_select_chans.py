# https://stackoverflow.com/questions/19130986/python-equivalent-of-golangs-select-on-channels

import queue
import os
import select
import threading
import time

class EQueue(queue.Queue):
    """
    TODO: super().put/get with os.eventfd_read/write(self._fd) MUST be atomic operation.
    """

    def __init__(self, *args, **kwargs):
        self._fd = os.eventfd(0, flags=os.EFD_CLOEXEC | os.EFD_NONBLOCK)
        super().__init__(*args, **kwargs)

    def put(self, *args, **kwargs):
        super().put(*args, **kwargs)
        os.eventfd_write(self._fd, 1)

    def get(self, *args, **kwargs):
        os.eventfd_read(self._fd)
        return super().get(*args, **kwargs)

    def fileno(self):
        return self._fd

    def __del__(self):
        os.close(self._fd)

def main():

    c1 = EQueue(maxsize=1)
    c2 = EQueue(maxsize=1)
    quit = EQueue(maxsize=1)

    def func1():
        for i in range(10):
            c1.put(i)

    threading.Thread(target=func1).start()

    def func2():
        for i in range(2):
            c2.put(i)

    threading.Thread(target=func2).start()

    while True:
        try:
            rx, _, _ = select.select([c1, c2, quit], [], [])
            if c1 in rx:
                print('Received value from c1', c1.get())
            if c2 in rx:
                print('Received value from c2', c2.get())
            if quit in rx:
                print('Received value from quit')
                break
        except KeyboardInterrupt:
            quit.put(0)

main()
