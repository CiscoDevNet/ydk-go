#!/bin/bash

if [[ $(uname) == "Linux" ]] && [[ -z "${GOPATH// }" ]]; then
    export GOPATH=/root/go-env
fi

echo "Installing YDK-Go"
go get github.com/CiscoDevNet/ydk-go/ydk

