#!/bin/bash

pkill sdm-master

sleep 1s

echo > ./logs/dm-master1.log

rm -rf ./data/dm-master1

./bin/sdm-master  --config=./conf/dm-master1.toml &   

