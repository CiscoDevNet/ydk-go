#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): dependencies_osx.sh | $@ ${NOCOLOR}"
}

function install_libssh {
    print_msg "Installing libssh-0.7.6"
    brew reinstall openssl
    export OPENSSL_ROOT_DIR=/usr/local/opt/openssl
    wget https://git.libssh.org/projects/libssh.git/snapshot/libssh-0.7.6.tar.gz
    tar zxf libssh-0.7.6.tar.gz && rm -f libssh-0.7.6.tar.gz
    mkdir libssh-0.7.6/build && cd libssh-0.7.6/build
    cmake ..
    sudo make install
    cd -
}

function install_golang {
    print_msg "Installing Go1.9.2"
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
    source /Users/travis/.gvm/scripts/gvm
    print_msg "GO version before installation: $(go version)"
    gvm install go1.4 -B
    gvm use go1.4
    export GOROOT_BOOTSTRAP=$GOROOT
    gvm install go1.9.2
    gvm use go1.9.2
    print_msg "GOROOT: $GOROOT"
    print_msg "GOPATH: $GOPATH"
    print_msg "GO version: $(go version)"
    print_msg " "
}

function install_libydk {
    print_msg "Installing YDK C++ core library"
    brew reinstall git
    git clone https://github.com/ciscodevnet/ydk-cpp.git -b 0.7.3
    mkdir ydk-cpp/core/ydk/build
    cd ydk-cpp/core/ydk/build
    cmake -DCMAKE_BUILD_TYPE=Release .. 
    sudo make install
    cd -

    #curl -O https://devhub.cisco.com/artifactory/osx-ydk/0.7.3/libydk-0.7.3-Darwin.pkg
    #sudo installer -pkg libydk-0.7.3-Darwin.pkg -target /
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

install_libssh

install_golang

install_libydk

brew rm -f --ignore-dependencies python python3
