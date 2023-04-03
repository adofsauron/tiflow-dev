#!/bin/bash

pkill sdm
pkill dm

sleep 1s


bin/sdm-master --name=sdm-192.168.75.149-8261 \
    --master-addr=0.0.0.0:8261 \
    --advertise-addr=192.168.75.149:8261 \
    --peer-urls=http://0.0.0.0:8291 \
    --advertise-peer-urls=http://192.168.75.149:8291 \
    --log-file=./logs/dm-master.log \
    --data-dir=./data/dm-master-8261 \
    --initial-cluster=sdm-192.168.75.149-8261=http://192.168.75.149:8291 \
    --config=./conf/dm-master1.toml &   


