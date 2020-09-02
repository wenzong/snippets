The Go Programming Language
========

## TL;DR

MacOS

```
$ brew install go
$ cat <<EOF >> ~/.bash_profile
export GOPATH=$HOME/workspace/go
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE=on
export GOPROXY=https://mirrors.aliyun.com/goproxy/ # Proxy
EOF
```
