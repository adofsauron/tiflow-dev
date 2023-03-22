# TiFlow

[![LICENSE](https://img.shields.io/github/license/pingcap/tiflow.svg)](https://sdbflow/blob/master/LICENSE)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/pingcap/tiflow)
![GitHub Release Date](https://img.shields.io/github/release-date/pingcap/tiflow)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pingcap/tiflow)
[![Build Status](https://sdbflow/actions/workflows/check_and_build.yaml/badge.svg?branch=master)](https://sdbflow/actions/workflows/check_and_build.yaml?query=event%3Apush+branch%3Amaster)
[![codecov](https://codecov.io/gh/pingcap/tiflow/branch/master/graph/badge.svg)](https://codecov.io/gh/pingcap/tiflow)
[![Coverage Status](https://coveralls.io/repos/github/pingcap/tiflow/badge.svg)](https://coveralls.io/github/pingcap/tiflow)
[![Go Report Card](https://goreportcard.com/badge/sdbflow)](https://goreportcard.com/report/sdbflow)

## Introduction

**TiFlow** is a unified data replication platform around [TiDB](https://docs.pingcap.com/tidb/stable),
including two main components:

* DM supports full data migration and incremental data replication from MySQL/MariaDB
  into [TiDB](https://docs.pingcap.com/tidb/stable).
* TiCDC supports replicating change data to various downstreams, including MySQL protocol-compatible databases
  and [Kafka](https://kafka.apache.org/).

More details can be found in [DM README](./README_DM.md) and [TiCDC README](./README_TiCDC.md).

## License

**TiFlow** is under the Apache 2.0 license. See the [LICENSE](./LICENSE) file for details.
