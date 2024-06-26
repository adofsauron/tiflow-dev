// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package conn

import (
	"context"

	"github.com/pingcap/errors"
	"github.com/pingcap/tidb-tools/pkg/dbutil"

	"sdbflow/dm/dm/config"
)

// FetchTimeZoneSetting fetch target db global time_zone setting.
func FetchTimeZoneSetting(ctx context.Context, cfg *config.DBConfig) (string, error) {
	db, err := DefaultDBProvider.Apply(cfg)
	if err != nil {
		return "", err
	}
	defer db.Close()
	dur, err := dbutil.GetTimeZoneOffset(ctx, db.DB)
	if err != nil {
		return "", errors.Trace(err)
	}
	return dbutil.FormatTimeZoneOffset(dur), nil
}
