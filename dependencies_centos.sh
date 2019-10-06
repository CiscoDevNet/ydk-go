#!/bin/bash

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

function check_install_gcc {
  which gcc
  local status=$?
  if [[ $status == 0 ]]
  then
    gcc_version=$(echo $(gcc --version) | awk '{ print $3 }')
    print_msg "Current gcc/g++ version is $gcc_version"
  else
    print_msg "The gcc/g++ not installed"
    gcc_version="4.0.0"
  fi
  if [[ $gcc_version < "4.8.1" ]]
  then
    print_msg "Upgrading gcc/g++ to version 5"
    yum install centos-release-scl -y > /dev/null
    yum install devtoolset-4-gcc* -y > /dev/null
    local status2=$?
    if [[ $status2 != 0 ]]; then
      MSG_COLOR=$RED
      print_msg "Failed to install gcc; exiting"
      exit 1
    else
      ln -sf /opt/rh/devtoolset-4/root/usr/bin/gcc /usr/bin/gcc
      ln -sf /opt/rh/devtoolset-4/root/usr/bin/g++ /usr/bin/g++

      gcc_version=$(echo $(gcc --version) | awk '{ print $3 }')
      print_msg "Installed gcc/g++ version is $gcc_version"
    fi
  fi
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

print_msg "Installing CentOS dependencies"
yum install -y epel-release
yum install -y git libxslt-devel libssh-devel gcc-c++ libstdc++-static cmake3 which wget make sudo > /dev/null

check_install_gcc

print_msg "Installing YDK C++ core library"
run_cmd yum install -y https://devhub.cisco.com/artifactory/rpm-ydk/0.8.4/libydk-0.8.4-1.x86_64.rpm > /dev/null

print_msg "Installing Golang version 1.9.2"
wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz &> /dev/null
tar -zxf  go1.9.2.linux-amd64.tar.gz -C /usr/local/
ln -sf /usr/local/go/bin/go /usr/bin/go
