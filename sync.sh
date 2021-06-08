#!/bin/bash

set +ex

servers=(
10.224.208.70
10.224.208.46
10.224.118.201
10.224.205.72
)

for server in ${servers[@]};
do
    scp grpc_mq root@$server:~/
    scp bootstrap.sh root@$server:~/
done