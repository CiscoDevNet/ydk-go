#!/bin/bash

function print_msg {
    echo -e "${MSG_COLOR}*** $(date): dependencies_trusty.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

print_msg "Installing Trusty OS dependencies"
sudo apt-get update > /dev/null
sudo apt-get install libssh-dev libtool gdebi-core -y > /dev/null

print_msg "Installing C++ version 5"
sudo apt-get install cmake gcc-5 g++-5 -y > /dev/null
sudo ln -f -s /usr/bin/g++-5 /usr/bin/c++
sudo ln -f -s /usr/bin/gcc-5 /usr/bin/cc

print_msg "Installing YDK C++ core library"
git clone https://github.com/ciscodevnet/ydk-cpp.git -b 0.8.1
mkdir ydk-cpp/core/ydk/build
cd ydk-cpp/core/ydk/build
cmake -DCMAKE_BUILD_TYPE=Release .. 
sudo make install
cd -

print_msg "Installing Golang version 1.9.2"
wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz &> /dev/null
sudo tar -zxf  go1.9.2.linux-amd64.tar.gz -C /usr/local/
sudo ln -f -s /usr/local/go/bin/go /usr/bin/go

