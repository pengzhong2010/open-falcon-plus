// Copyright 2017 Xiaomi, Inc.
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

package api

import (
	"testing"

	"github.com/pengzhong2010/open-falcon-plus/falcon-plus/modules/alarm/g"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	g.ParseConfig("../cfg.example.json")
}

func TestPortalAPI(t *testing.T) {
	Convey("Get action from api failed", t, func() {
		r := CurlAction(1)
		So(r.Id, ShouldEqual, 1)
	})

}
