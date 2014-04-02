import signal
import sys
import types

def sig_handler(signal, frame):
    print(signal)
    print(isinstance(frame, types.FrameType) )
    sys.exit(0)

signal.signal(signal.SIGINT, sig_handler)

signal.pause()
