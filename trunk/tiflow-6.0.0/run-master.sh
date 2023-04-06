#!/bin/bash

pkill sdm-master

sleep 1s

rm -rf ./data/dm-master1

touch ./logs/dm-master1.log

./bin/sdm-master  --config=./conf/dm-master1.toml &   

