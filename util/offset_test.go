package util

import (
	"testing"
)

func TestOffset(t *testing.T) {
	testcases := []struct {
		name        string
		page        int
		pageSize    int
		defPageSize int
		expect      int
	}{
		{
			"given page is negative",
			-1,
			-1,
			10,
			0,
		},
		{
			"given page is 0",
			0,
			-1,
			10,
			0,
		},
		{
			"given page is 1",
			1,
			5,
			10,
			0,
		},
		{
			"given page is greater than 1",
			2,
			5,
			10,
			5,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			out := Offset(tt.page, tt.pageSize, tt.defPageSize)
			if out != tt.expect {
				t.Errorf("want %d, got %d", tt.expect, out)
			}
		})
	}
}
