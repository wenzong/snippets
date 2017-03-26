import getpass
import os
import sys

import paramiko


def lookup_ssh_config(machine):
    ssh_config = paramiko.config.SSHConfig()
    _f = os.path.expanduser('~/.ssh/config')
    if os.path.isfile(_f):
        with open(_f, 'r') as f:
            ssh_config.parse(f)
    return ssh_config.lookup(machine)

host = sys.argv[1]

host_info = lookup_ssh_config(host)
dst_hostname = host_info.get('hostname')
identityfiles = host_info.get('identityfile')
if len(identityfiles) > 0:
    _f = os.path.expanduser(identityfiles[0])
    private_key = paramiko.RSAKey(filename=_f)
else:
    private_key = None
username = host_info.get('user', getpass.getuser())
port = int(host_info.get('port', 22))


client = paramiko.SSHClient()
client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
client.load_system_host_keys()
if private_key:
    client.connect(dst_hostname, username=username, port=port, pkey=private_key)
else:
    client.connect(dst_hostname, username=username, port=port)

stdin, stdout, stderr = client.exec_command('ls')
for line in stdout:
    print(line.strip('\n'))
client.close()
