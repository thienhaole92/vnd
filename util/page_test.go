package util

import "testing"

func TestPage(t *testing.T) {
	testcases := []struct {
		name   string
		value  int
		expect int
	}{
		{
			"given value is less than 0",
			-1,
			1,
		},
		{
			"given value is equal 0",
			0,
			1,
		},
		{
			"given value is greater than 0",
			10,
			10,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			out := Page(tt.value)
			if out != tt.expect {
				t.Errorf("want %d, got %d", tt.expect, out)
			}
		})
	}
}
