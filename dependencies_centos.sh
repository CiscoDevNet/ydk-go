#!/bin/sh

yum install -y epel-release
yum install -y go git libxslt-devel libssh-devel libstdc++-static python-devel cmake3 python-pip which make sudo > /dev/null
yum install -y https://devhub.cisco.com/artifactory/rpm-ydk/0.7.2/libydk-0.7.2-1.x86_64.rpm > /dev/null

yum install centos-release-scl -y > /dev/null
yum install devtoolset-4-gcc* -y > /dev/null

ln -sf /opt/rh/devtoolset-4/root/usr/bin/gcc /usr/bin/gcc
ln -sf /opt/rh/devtoolset-4/root/usr/bin/g++ /usr/bin/g++

sudo easy_install pip
