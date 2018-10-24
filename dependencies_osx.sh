#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): dependencies_osx.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

print_msg "Installing libssh 7.5.5"
brew install openssl
export OPENSSL_ROOT_DIR=/usr/local/opt/openssl
wget https://git.libssh.org/projects/libssh.git/snapshot/libssh-0.7.6.tar.gz
tar zxf libssh-0.7.6.tar.gz
cd libssh-0.7.6
mkdir build && cd build
cmake ..
sudo make install

brew rm -f --ignore-dependencies python python3

print_msg "Installing YDK C++ core library"
curl -O https://devhub.cisco.com/artifactory/osx-ydk/0.7.3/libydk-0.7.3-Darwin.pkg
sudo installer -pkg libydk*pkg -target /
