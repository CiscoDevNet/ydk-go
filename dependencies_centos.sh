#!/bin/sh

function print_msg {
    echo -e "${MSG_COLOR}*** $(date): dependencies_centos.sh | $@ ${NOCOLOR}"
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

print_msg "Installing CentOS dependencies"
yum install -y epel-release
yum install -y git libxslt-devel libssh-devel libstdc++-static python-devel cmake3 python-pip which wget make sudo > /dev/null

print_msg "Installing YDK 0.7.3 core library"
run_cmd yum install -y https://devhub.cisco.com/artifactory/rpm-ydk/0.7.3/libydk-0.7.3-1.x86_64.rpm > /dev/null

print_msg "Installing C++ version 5"
yum install centos-release-scl -y > /dev/null
yum install devtoolset-4-gcc* -y > /dev/null

ln -sf /opt/rh/devtoolset-4/root/usr/bin/gcc /usr/bin/gcc
ln -sf /opt/rh/devtoolset-4/root/usr/bin/g++ /usr/bin/g++

print_msg "Installing Golang version 1.9.2"
wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz &> /dev/null
tar -zxf  go1.9.2.linux-amd64.tar.gz -C /usr/local/
ln -s /usr/local/go/bin/go /usr/bin/go

sudo easy_install pip
