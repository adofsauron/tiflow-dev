#!/bin/bash

systemctl stop grafana-server

sleep 3s

systemctl start grafana-server

