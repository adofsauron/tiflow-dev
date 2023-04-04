#!/bin/bash

pkill sdm-master

sleep 1s


./bin/sdm-master  --config=./conf/dm-master1.toml &   

