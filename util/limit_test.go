package util

import "testing"

func TestLimit(t *testing.T) {
	testcases := []struct {
		name       string
		value      int
		defaultVal int
		expect     int
	}{
		{
			"value less than 0, default less than 0",
			-1,
			-1,
			DEFAULT_PAGE_ITEMS,
		},
		{
			"value equal 0, default equal 0",
			0,
			0,
			DEFAULT_PAGE_ITEMS,
		},

		{
			"value less than 0, default greater than 0",
			-1,
			10,
			10,
		},
		{
			"value greater than 0, default greater than 0",
			10,
			10,
			10,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			out := Limit(tt.value, tt.defaultVal)
			if out != tt.expect {
				t.Errorf("want %d, got %d", tt.expect, out)
			}
		})
	}
}
