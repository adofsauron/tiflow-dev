// Copyright 2022 PingCAP, Inc.
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

package partition

import (
	"sdbflow/cdc/model"
	"sdbflow/pkg/hash"
)

// TableDispatcher is a partition dispatcher which dispatches events
// based on the schema and table name.
type TableDispatcher struct {
	hasher *hash.PositionInertia
}

// NewTableDispatcher creates a TableDispatcher.
func NewTableDispatcher() *TableDispatcher {
	return &TableDispatcher{
		hasher: hash.NewPositionInertia(),
	}
}

// DispatchRowChangedEvent returns the target partition to which
// a row changed event should be dispatched.
func (t *TableDispatcher) DispatchRowChangedEvent(row *model.RowChangedEvent, partitionNum int32) int32 {
	t.hasher.Reset()
	// distribute partition by table
	t.hasher.Write([]byte(row.Table.Schema), []byte(row.Table.Table))
	return int32(t.hasher.Sum32() % uint32(partitionNum))
}