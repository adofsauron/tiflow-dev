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

package utils

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-mysql-org/go-mysql/replication"
	"sdbflow/dm/pkg/terror"
)

// not support to config yet.
var (
	UUIDIndexFilename  = "server-uuid.index"
	MetaFilename       = "relay.meta"
	uuidIndexSeparator = "."
)

// ParseUUIDIndex parses UUIDIndexFilename, return a list of relay log subdirectory names and error.
func ParseUUIDIndex(indexPath string) ([]string, error) {
	fd, err := os.Open(indexPath)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, terror.ErrRelayParseUUIDIndex.Delegate(err)
	}
	defer fd.Close()

	uuids := make([]string, 0, 5)
	br := bufio.NewReader(fd)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			if len(line) == 0 {
				break
			}
		} else if err != nil {
			return nil, terror.ErrRelayParseUUIDIndex.Delegate(err)
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		uuids = append(uuids, line)
	}

	return uuids, nil
}

// AddSuffixForUUID adds a suffix for UUID, returns the name for relay log subdirectory.
func AddSuffixForUUID(uuid string, id int) string {
	return fmt.Sprintf("%s%s%06d", uuid, uuidIndexSeparator, id) // eg. 53ea0ed1-9bf8-11e6-8bea-64006a897c73.000001
}

// SuffixIntToStr convert int-represented suffix to string-represented.
// TODO: assign RelaySubDirSuffix a type and implement Stringer.
func SuffixIntToStr(id int) string {
	return fmt.Sprintf("%06d", id)
}

// ParseRelaySubDir parses relay log subdirectory name to (server UUID, RelaySubDirSuffix) pair.
func ParseRelaySubDir(uuid string) (string, int, error) {
	parts := strings.Split(uuid, uuidIndexSeparator)
	if len(parts) != 2 || len(parts[1]) != 6 {
		return "", 0, terror.ErrRelayParseUUIDSuffix.Generate(uuid)
	}
	ID, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, terror.ErrRelayParseUUIDSuffix.Generate(uuid)
	}
	return parts[0], ID, nil
}

// GetUUIDBySuffix gets relay log subdirectory name by matching suffix.
func GetUUIDBySuffix(uuids []string, suffix string) string {
	suffix2 := fmt.Sprintf("%s%s", uuidIndexSeparator, suffix)
	for _, uuid := range uuids {
		if strings.HasSuffix(uuid, suffix2) {
			return uuid
		}
	}
	return ""
}

// GenFakeRotateEvent generates a fake ROTATE_EVENT without checksum
// ref: https://github.com/mysql/mysql-server/blob/4f1d7cf5fcb11a3f84cff27e37100d7295e7d5ca/sql/rpl_binlog_sender.cc#L855
func GenFakeRotateEvent(nextLogName string, logPos uint64, serverID uint32) (*replication.BinlogEvent, error) {
	headerLen := replication.EventHeaderSize
	bodyLen := 8 + len(nextLogName)
	eventSize := headerLen + bodyLen

	rawData := make([]byte, eventSize)

	// header
	binary.LittleEndian.PutUint32(rawData, 0)                         // timestamp
	rawData[4] = byte(replication.ROTATE_EVENT)                       // event type
	binary.LittleEndian.PutUint32(rawData[4+1:], serverID)            // server ID
	binary.LittleEndian.PutUint32(rawData[4+1+4:], uint32(eventSize)) // event size
	binary.LittleEndian.PutUint32(rawData[4+1+4+4:], 0)               // log pos, always 0
	binary.LittleEndian.PutUint16(rawData[4+1+4+4+4:], 0x20)          // flags, LOG_EVENT_ARTIFICIAL_F

	// body
	binary.LittleEndian.PutUint64(rawData[headerLen:], logPos)
	copy(rawData[headerLen+8:], nextLogName)

	// decode header
	h := &replication.EventHeader{}
	err := h.Decode(rawData)
	if err != nil {
		return nil, terror.ErrRelayGenFakeRotateEvent.Delegate(err)
	}

	// decode body
	e := &replication.RotateEvent{}
	err = e.Decode(rawData[headerLen:])
	if err != nil {
		return nil, terror.ErrRelayGenFakeRotateEvent.Delegate(err)
	}

	return &replication.BinlogEvent{RawData: rawData, Header: h, Event: e}, nil
}
