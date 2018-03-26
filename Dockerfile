FROM ubuntu:xenial

MAINTAINER http://ydk.io

COPY . /root/ydk-go

RUN echo 'Installing dependencies and ydk-go'

WORKDIR /root/ydk-go

RUN /bin/bash -c './dependencies_xenial.sh && ./install.sh'
