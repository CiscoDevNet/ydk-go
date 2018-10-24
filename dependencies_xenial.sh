#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): dependencies_xenial.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

print_msg "Installing Xenial OS dependencies"
apt update -y > /dev/null
apt install sudo -y > /dev/null
apt-get update > /dev/null
sudo apt-get install libtool-bin -y > /dev/null
local status=$?
if [[ ${status} != 0 ]]; then
   sudo apt-get install libtool -y > /dev/null
fi

apt-get install git gdebi-core python-dev python-pip wget -y > /dev/null

print_msg "Installing C++ version 5"
add-apt-repository ppa:ubuntu-toolchain-r/test -y
apt-get install gcc-5 g++-5 -y > /dev/null
ln -f -s /usr/bin/g++-5 /usr/bin/c++
ln -f -s /usr/bin/gcc-5 /usr/bin/cc

print_msg "Installing Golang version 1.9.2"
wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz &> /dev/null
tar -zxf  go1.9.2.linux-amd64.tar.gz -C /usr/local/
ln -f -s /usr/local/go/bin/go /usr/bin/go

print_msg "Installing YDK C++ core library"
wget https://devhub.cisco.com/artifactory/debian-ydk/0.7.3/libydk_0.7.3-1_amd64.deb
gdebi -n libydk_0.7.3-1_amd64.deb

print_msg "Installing pip"
sudo easy_install pip
