#!/bin/bash

if [[ $# != 1 ]];then
    echo "USAGE:bash $0 SERVER_TYPE<grpc||http>"
    echo "eg:bash $0 grpc"
fi

server_type=$1
if [[ $1 == "grpc" ]];then
    ./grpc_mq --transport=grpc ｜ grep -v 10.28.23.217
elif [[ $1 == "http" ]];then
    ./grpc_mq | grep -v 10.28.23.217
fi