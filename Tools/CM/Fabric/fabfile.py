from fabric.connection import Connection

with Connection('pi') as c:
    c.run('uname -a')
