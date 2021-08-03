FROM ubuntu:latest
WORKDIR ./go/akina
COPY ./bin/akina ./
COPY ./akina.db ./

RUN apt-get update
RUN DEBIAN_FRONTEND="noninteractive" apt-get -y install tzdata
RUN apt-get install -y ca-certificates

CMD ["./akina"]