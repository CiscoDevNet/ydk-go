#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): dependencies_trusty.sh | $@ ${NOCOLOR}"
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

print_msg "Installing Trusty OS dependencies"
sudo apt-get update > /dev/null
sudo apt-get install libssh-dev libtool gdebi-core python3-dev python-dev -y > /dev/null
sudo add-apt-repository ppa:ubuntu-toolchain-r/test -y > /dev/null
sudo apt-get update > /dev/null

print_msg "Installing C++ version 5"
sudo apt-get install cmake gcc-5 g++-5 -y > /dev/null
sudo ln -f -s /usr/bin/g++-5 /usr/bin/c++
sudo ln -f -s /usr/bin/gcc-5 /usr/bin/cc

print_msg "Installing YDK 0.7.3 C++ core library"
git clone https://github.com/ciscodevnet/ydk-cpp.git -b 0.7.3
mkdir ydk-cpp/core/ydk/build
cd ydk-cpp/core/ydk/build
cmake -DCMAKE_BUILD_TYPE=Release .. 
sudo make install
cd -

sudo easy_install pip
sudo pip install pybind11
