// Copyright 2019 PingCAP, Inc.
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

package main

// Reference: https://dzone.com/articles/measuring-integration-test-coverage-rate-in-pouchc

//nolint: gofumpt
import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	_ "sdbflow/dm/pkg/printinit"
)

func TestRunMain(_ *testing.T) {
	fmt.Println("dm-master startup", time.Now())
	var args []string
	for _, arg := range os.Args {
		switch {
		case arg == "DEVEL":
		case strings.HasPrefix(arg, "-test."):
		default:
			args = append(args, arg)
		}
	}

	os.Args = args
	main()
	fmt.Println("dm-master exit", time.Now())
}
