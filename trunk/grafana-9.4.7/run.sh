#!/bin/bash


pkill grafana

sleep 2s

./bin/linux-amd64/grafana-server  --config=./conf/defaults.ini  web &

