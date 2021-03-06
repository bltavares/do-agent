// Copyright 2016 DigitalOcean
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package procfs

import (
	"strings"
	"testing"
)

const testOSReleaseValues = `3.13.0-71-generic`

func TestOSRelease(t *testing.T) {
	r, err := readOSRelease(strings.NewReader(testOSReleaseValues))
	if err != nil {
		t.Errorf("Unable to read test values")
	}

	expectedRelease := OSRelease("3.13.0-71-generic")
	if r != expectedRelease {
		t.Errorf("OSRelease: expected %s, actual %s", expectedRelease, r)
	}
}
