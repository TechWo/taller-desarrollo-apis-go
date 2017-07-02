# Dockerfile for building a development environment to work with
# Golang to create RESTFUL APIs.

FROM ubuntu:16.04
MAINTAINER Marcela Bonell <mbonell@techwomen.org.mx>

# Update the repository sources list and install dependencies
RUN apt-get update && apt-get install -y \
  vim \
  git \
  curl

# Installing Golang environment
RUN curl -O https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz

ENV GOPATH=/go-workspace
ENV PATH=$PATH:/usr/local/go/bin

WORKDIR $GOPATH

# Setting up Golang workspace
RUN mkdir bin \
    && mkdir pkg \
    && mkdir src

# Installing dependencies
RUN cd src \
  && go get github.com/gorilla/mux \
  && go get github.com/alecthomas/gometalinter


# Installing vim-go as IDE
RUN git clone https://github.com/farazdagi/vim-go-ide.git ~/.vim_go_runtime && sh ~/.vim_go_runtime/bin/install \
    && mkdir -p ~/.vim/autoload ~/.vim/bundle && curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim \
    && echo 'alias vimgo="vim -u ~/.vimrc.go"' >> ~/.bashrc \
    && echo 'set t_Co=256' >> ~/.vimrc.go
