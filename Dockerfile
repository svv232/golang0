FROM ubuntu:16.04
MAINTAINER sai
LABEL Description="A chat platform by Sai Vegasena"

#installation 
RUN apt-get clean 
RUN apt-get update && apt-get upgrade -y
RUN apt-get install golang

#user 
WORKDIR /app
