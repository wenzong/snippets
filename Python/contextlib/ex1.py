"""
https://www.python.org/dev/peps/pep-0343/

https://mail.python.org/pipermail/python-dev/2005-July/054658.html
"""

from contextlib import contextmanager

@contextmanager
def tag(name):
    print('<%s>' % name)
    yield
    print('</%s>' % name)


with tag('h1'):
    print('foo')
