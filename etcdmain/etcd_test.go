// Copyright 2015 CoreOS, Inc.
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

package etcdmain

import (
	"testing"

	"github.com/coreos/etcd/pkg/testutil"
)

func TestGenClusterString(t *testing.T) {
	tests := []struct {
		token string
		urls  []string
		wstr  string
	}{
		{
			"default", []string{"http://127.0.0.1:2379"},
			"default=http://127.0.0.1:2379",
		},
		{
			"node1", []string{"http://0.0.0.0:2379", "http://1.1.1.1:2379"},
			"node1=http://0.0.0.0:2379,node1=http://1.1.1.1:2379",
		},
	}
	for i, tt := range tests {
		urls := testutil.MustNewURLs(t, tt.urls)
		str := genClusterString(tt.token, urls)
		if str != tt.wstr {
			t.Errorf("#%d: cluster = %s, want %s", i, str, tt.wstr)
		}
	}
}
