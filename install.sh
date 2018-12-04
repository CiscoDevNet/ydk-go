#!/bin/bash

function print_msg {
    echo -e "$MSG_COLOR*** $(date): install.sh | $@ $NOCOLOR"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

if [[ $(uname) == "Darwin" ]]; then
    print_msg "Installing go1.9.2"
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer) 
    source $HOME/.gvm/scripts/gvm
    print_msg "GO version before installation: $(go version)"
    gvm install go1.4 -B
    gvm use go1.4
    export GOROOT_BOOTSTRAP=$GOROOT
    gvm install go1.9.2
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