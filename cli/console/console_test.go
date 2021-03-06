// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package console

import (
	"bytes"
	"strings"
	"testing"
)

// fakeDevice inherits all members of target.Device through embedding.
// Unimplemented members send a clear signal during tests because they will
// panic if called, allowing us to implement only the minimum set of members
// required for testing.
type fakeDevice struct {
	id           string
	friendlyName string
	size         uint64
}

func (f *fakeDevice) Identifier() string {
	return f.id
}

func (f *fakeDevice) FriendlyName() string {
	return f.friendlyName
}

func (f *fakeDevice) Size() uint64 {
	return f.size
}

func TestPrintDevices(t *testing.T) {
	deviceOne := &fakeDevice{
		id:           "drive1",
		friendlyName: "foo super duper drive",
		size:         1123456789,
	}
	deviceTwo := &fakeDevice{
		id:           "drive2",
		friendlyName: "bar bodacious drive",
		size:         9987654321,
	}
	deviceThree := &fakeDevice{
		id:           "drive3",
		friendlyName: "baz radical drive",
		size:         19987654321,
	}
	tests := []struct {
		desc    string
		devices []TargetDevice
	}{
		{
			desc:    "no devices",
			devices: []TargetDevice{},
		},
		{
			desc:    "one device",
			devices: []TargetDevice{deviceOne},
		},
		{
			desc:    "two devices",
			devices: []TargetDevice{deviceOne, deviceTwo},
		},
		{
			desc:    "three devices",
			devices: []TargetDevice{deviceOne, deviceTwo, deviceThree},
		},
	}
	for _, tt := range tests {
		var got bytes.Buffer
		PrintDevices(tt.devices, &got)
		for _, device := range tt.devices {
			if !strings.Contains(got.String(), device.Identifier()) {
				t.Errorf("%s: PrintDevices() got = %q, must contain = %q", tt.desc, got.String(), device.Identifier())
			}
		}
	}
}
