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

package cli

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pingcap/errors"
	"sdbflow/cdc/model"
	"sdbflow/pkg/api/v1/mock"
	cerror "sdbflow/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestChangefeedRemoveCli(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cf := mock.NewMockChangefeedInterface(ctrl)
	f := &mockFactory{changefeeds: cf}

	cmd := newCmdRemoveChangefeed(f)

	cf.EXPECT().Get(gomock.Any(), "abc").Return(&model.ChangefeedDetail{}, nil)
	cf.EXPECT().Delete(gomock.Any(), "abc").Return(nil)
	cf.EXPECT().Get(gomock.Any(), "abc").Return(nil,
		cerror.ErrChangeFeedNotExists.GenWithStackByArgs("abc"))
	os.Args = []string{"remove", "--changefeed-id=abc"}
	require.Nil(t, cmd.Execute())
	cf.EXPECT().Get(gomock.Any(), "abc").Return(nil,
		cerror.ErrChangeFeedNotExists.GenWithStackByArgs("abc"))
	os.Args = []string{"remove", "--changefeed-id=abc"}
	require.Nil(t, cmd.Execute())

	o := newRemoveChangefeedOptions()
	o.complete(f)
	o.changefeedID = "abc"
	cf.EXPECT().Get(gomock.Any(), "abc").Return(nil, errors.New("abc"))
	require.NotNil(t, o.run(cmd))
}
