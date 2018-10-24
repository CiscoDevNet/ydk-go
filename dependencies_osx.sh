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
brew reinstall openssl
export OPENSSL_ROOT_DIR=/usr/local/opt/openssl
wget https://git.libssh.org/projects/libssh.git/snapshot/libssh-0.7.6.tar.gz
tar zxf libssh-0.7.6.tar.gz
cd libssh-0.7.6
mkdir build && cd build
cmake ..
sudo make install
cd ../../

brew rm -f --ignore-dependencies python python3

print_msg "Installing Go1.9.2"
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.9.2
gvm use go1.9.2
print_msg "GOROOT: $GOROOT"
print_msg "GOPATH: $GOPATH"
print_msg "GO version: $(go version)"
print_msg " "

print_msg "Installing YDK C++ core library"
curl -O https://devhub.cisco.com/artifactory/osx-ydk/0.7.3/libydk-0.7.3-Darwin.pkg
sudo installer -pkg libydk*pkg -target /
