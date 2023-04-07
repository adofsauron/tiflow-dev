#!/bin/bash

./bin/sdmctl --master-addr=192.168.30.40:8261 operate-source create ./conf/mysql-01.yaml


./bin/sdmctl --master-addr=192.168.30.40:8261 start-task ./conf/testdm-task.yaml


./bin/sdmctl --master-addr=192.168.30.40:8261 query-status test


./bin/sdmctl --master-addr=192.168.30.40:8261 operate-source stop ./conf/mysql-01.yaml


./bin/sdmctl --master-addr=192.168.30.40:8261 stop-task ./conf/testdm-task.yaml


