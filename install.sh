#!/bin/bash

function print_msg {
    echo -e "${MSG_COLOR}*** $(date): install.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

if [[ $(uname) == "Darwin" ]]; then
    source /Users/travis/.gvm/scripts/gvm
    gvm use go1.9.2
    print_msg "GOROOT: $GOROOT"
    print_msg "GOPATH: $GOPATH"
    print_msg "GO version: $(go version)"
else
    export GOPATH="$(pwd)/golang"
    export GOROOT=/usr/local/go
    export PATH=$GOROOT/bin:$PATH
    print_msg "Setting GOROOT to $GOROOT"
    print_msg "Setting GOPATH to $GOPATH"
fi

print_msg "Installing YDK-Go core package"
go get github.com/CiscoDevNet/ydk-go/ydk
