#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): install.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

print_msg "Installing OSX dependencies"
brew update > /dev/null
brew install libssh xml2 curl pybind11 > /dev/null

brew rm -f --ignore-dependencies python python3

print_msg "Installing YDK 0.7.3 C++ core library"
curl -O https://devhub.cisco.com/artifactory/osx-ydk/0.7.3/libydk-0.7.3-Darwin.pkg
sudo installer -pkg libydk*pkg -target /
