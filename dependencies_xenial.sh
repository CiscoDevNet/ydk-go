#!/bin/sh

function print_msg {
    echo -e "${MSG_COLOR}*** $(date): dependencies_xenial.sh | $@ ${NOCOLOR}"
}

function run_cmd {
    print_msg "Running command: $@"
    $@
    local status=$?
    if [ $status -ne 0 ]; then
        MSG_COLOR=$RED
        print_msg "Exiting '$@' with status=$status"
        exit $status
    fi
    return $status
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

print_msg "Installing Xenial OS dependencies"
apt-get update -y > /dev/null
apt-get install golang git gdebi-core python-dev python-pip libtool-bin wget sudo -y > /dev/null
add-apt-repository ppa:ubuntu-toolchain-r/test -y
apt-get update > /dev/null

print_msg "Installing C++ version 5"
apt-get install gcc-5 g++-5 -y > /dev/null
ln -f -s /usr/bin/g++-5 /usr/bin/c++
ln -f -s /usr/bin/gcc-5 /usr/bin/cc

print_msg "Installing YDK 0.7.3 core library"
run_cmd wget https://devhub.cisco.com/artifactory/debian-ydk/0.7.3/libydk_0.7.3-1_amd64.deb
gdebi -n libydk_0.7.3-1_amd64.deb

print_msg "Installing Golang version 1.9.2"
wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz &> /dev/null
tar -zxf  go1.9.2.linux-amd64.tar.gz -C /usr/local/
ln -s /usr/local/go/bin/go /usr/bin/go

sudo easy_install pip
