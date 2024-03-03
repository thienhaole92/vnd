package share

import (
	"testing"
)

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
			DefaultPageItems,
		},
		{
			"value equal 0, default equal 0",
			0,
			0,
			DefaultPageItems,
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
