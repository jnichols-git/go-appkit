// Copyright 2023 appkit Authors
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

package output

import (
	"testing"

	"github.com/decentplatforms/appkit/logf"
	"github.com/decentplatforms/appkit/logf/formats"
)

func TestFile(t *testing.T) {
	writer, err := Open("../cicd/test.log", 100)
	if err != nil {
		t.Fatal(err)
	}
	log, err := logf.NewLogger(logf.Config{
		MaxLevel:     logf.Debug,
		DefaultLevel: logf.Informational,
		Format:       formats.Syslog5424Format(formats.SyslogConfig{Tag: "file-test"}),
		Output:       writer,
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Log(logf.Informational, "test log")
}
