# This Dockerfile defines the image which jeeves is executed in

# Repository build command: sudo docker build -t="atomos/default:latest" .

FROM ubuntu:14.04
MAINTAINER Atomos <support@atomos.io>

# Update currently installed packages
# Install R, Python, Curl, GCC/G++, Perl (Just to make sure), and the SciPy stack (not including Ipython Notebook)
# Install Go
# Install Java
# Download jeeves into the /usr/local/bin folder and set permissions

RUN apt-get update && apt-get install r-base r-base-dev python curl build-essential perl python-numpy python-scipy python-matplotlib ipython python-pandas python-sympy python-nose -y && curl -O https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.3.1.linux-amd64.tar.gz && rm go1.3.1.linux-amd64.tar.gz && export PATH=$PATH:/usr/local/go/bin && curl -O http://storage.googleapis.com/assets.atomos.io/carbon/jdk-8u20-linux-x64.tar.gz && tar -C /usr/local -xzf jdk-8u20-linux-x64.tar.gz && rm jdk-8u20-linux-x64.tar.gz && export PATH=$PATH:/usr/local/jdk1.8.0_20/bin && curl https://raw.githubusercontent.com/AtomosIO/jeeves/master/jeeves -o /usr/local/bin/jeeves && chmod 750 /usr/local/bin/jeeves
