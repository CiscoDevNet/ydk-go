#!/bin/bash

brew update > /dev/null
brew install libssh xml2 curl pybind11 > /dev/null

brew rm -f --ignore-dependencies python python3

curl -O https://devhub.cisco.com/artifactory/osx-ydk/0.8.0/libydk-0.8.0-Darwin.pkg
sudo installer -pkg libydk*pkg -target /

