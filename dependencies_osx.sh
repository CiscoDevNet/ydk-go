#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): dependencies_osx.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

print_msg "Installing OSX dependencies"
brew update > /dev/null
brew install libssh xml2 curl > /dev/null

brew rm -f --ignore-dependencies python python3

print_msg "Installing Golang version 1.9.2"
export CGO_ENABLED=0
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.9.2

print_msg "Installing YDK C++ core library"
curl -O https://devhub.cisco.com/artifactory/osx-ydk/0.7.3/libydk-0.7.3-Darwin.pkg
sudo installer -pkg libydk*pkg -target /
