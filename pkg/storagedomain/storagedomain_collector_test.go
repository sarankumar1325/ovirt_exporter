// SPDX-License-Identifier: MIT

package storagedomain

import "testing"

func TestStorageUpValue(t *testing.T) {
	cases := []struct {
		status string
		want   float64
	}{
		{status: "active", want: 1},
		{status: "maintenance", want: 2},
		{status: "preparing_for_maintenance", want: 2},
		{status: "unattached", want: 3},
		{status: "detached", want: 3},
		{status: "inactive", want: 0},
		{status: "locked", want: 0},
		{status: "unknown", want: 0},
		{status: "", want: 0},
	}
	for _, tc := range cases {
		if got := storageUpValue(tc.status); got != tc.want {
			t.Errorf("storageUpValue(%q) = %v, want %v", tc.status, got, tc.want)
		}
	}
}
