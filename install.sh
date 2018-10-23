#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): install.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

if [[ $(uname) == "Linux" ]] && [[ -z "${GOPATH// }" ]]; then
    export GOPATH=/home/ygorelik/go-env
    print_msg "Setting GOPATH to $GOPATH"
fi

print_msg "Installing YDK-Go core package"
go get github.com/CiscoDevNet/ydk-go/ydk
