# 配置博客站点
FROM ubuntu
MAINTAINER Berkaroad "jiarong.bai198605@gmail.com"
RUN apt-get update
RUN apt-get install -y openssh-server
RUN apt-get install -y golang
RUN apt-get install -y mongodb