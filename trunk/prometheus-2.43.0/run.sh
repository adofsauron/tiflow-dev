#!/bin/bash

pkill prometheus

sleep 3s 

exec > >(tee -i -a "./logs/prometheus.log")
exec 2>&1

exec ./prometheus \
    --config.file="./conf/prometheus.yml" \
    --web.listen-address=":9090" \
    --web.external-url="http://192.168.30.40:9090/" \
    --web.enable-admin-api \
    --log.level="info" \
    --storage.tsdb.path="./data/" \
    --storage.tsdb.retention="30d" &





