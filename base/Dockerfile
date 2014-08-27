# This Dockerfile defines the image which jeeves is executed in

# Repository build command: sudo docker build -t="atomos/default:latest" .

FROM ubuntu:14.04
MAINTAINER Atomos <support@atomos.io>

# Update currently installed packages
RUN apt-get update

# Install R, Python, Ruby, Curl, GCC/G++, and Perl (Just to make sure) 
RUN apt-get install r-base r-base-dev python ruby curl build-essential perl -y

# Install Go
RUN curl -O https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.3.1.linux-amd64.tar.gz
RUN rm go1.3.1.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin

# Install Java
RUN curl -O http://storage.googleapis.com/assets.atomos.io/carbon/jdk-8u20-linux-x64.tar.gz
RUN tar -C /usr/local -xzf jdk-8u20-linux-x64.tar.gz
RUN rm jdk-8u20-linux-x64.tar.gz
RUN export PATH=$PATH:/usr/local/jdk1.8.0_20/bin

# Download jeeves into the /usr/local/bin folder and set permissions
RUN curl https://raw.githubusercontent.com/AtomosIO/jeeves/master/jeeves -o /usr/local/bin/jeeves
RUN chmod 750 /usr/local/bin/jeeves