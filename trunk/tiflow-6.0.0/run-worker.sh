#!/bin/bash

pkill sdm-worker

sleep 1s

./bin/sdm-worker --config=./conf/dm-worker1.toml & 