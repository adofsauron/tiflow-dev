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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// directio is a proxy package for github.com/pingcap/badger/directio
package directio

import (
	"github.com/ncw/directio"
)

const (
	// AlignSize is the size to align the buffer to
	AlignSize = directio.AlignSize
	// BlockSize is the minimum block size
	BlockSize = directio.BlockSize
)

// AlignedBlock returns []byte of size BlockSize aligned to a multiple
// of AlignSize in memory (must be power of two)
var AlignedBlock = directio.AlignedBlock
