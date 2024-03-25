package util

import "testing"

func TestBuildConitionalQuery(t *testing.T) {
	type Expectation struct {
		Where string
		Args  []any
	}

	id, fn, ln := 123, "John", "Doe"

	testcases := []struct {
		name      string
		arg       any
		startFrom uint
		expect    Expectation
	}{
		{
			"no column tag specified",
			struct {
				Offset int
				Limit  int
			}{
				Offset: 1,
				Limit:  1,
			},
			1,
			Expectation{
				Where: "",
				Args:  []any{},
			},
		},
		{
			"all column tags specified start from 1",
			struct {
				Id        *int    `col:"id"`
				FirstName *string `col:"first_name"`
				LastName  *string `col:"last_name"`
			}{
				Id:        &id,
				FirstName: &fn,
				LastName:  &ln,
			},
			1,
			Expectation{
				Where: "WHERE id = $1 AND first_name = $2 AND last_name = $3",
				Args:  []any{123, "John", "Doe"},
			},
		},
		{
			"some column tags specified start from 1",
			struct {
				Id        *int    `col:"id"`
				FirstName *string `col:"first_name"`
				LastName  *string `col:"last_name"`
			}{
				Id:        nil,
				FirstName: &fn,
				LastName:  &ln,
			},
			1,
			Expectation{
				Where: "WHERE first_name = $1 AND last_name = $2",
				Args:  []any{"John", "Doe"},
			},
		},
		{
			"mixed fields with and without column tags start from 1",
			struct {
				Id        int
				FirstName *string `col:"first_name"`
				LastName  *string `col:"last_name"`
			}{
				Id:        1,
				FirstName: &fn,
				LastName:  &ln,
			},
			1,
			Expectation{
				Where: "WHERE first_name = $1 AND last_name = $2",
				Args:  []any{"John", "Doe"},
			},
		},
		{
			"mixed pointer fields and non pointer colufieldsmn tags start from 1",
			struct {
				Id        int     `col:"id"`
				FirstName *string `col:"first_name"`
				LastName  *string `col:"last_name"`
			}{
				Id:        1,
				FirstName: &fn,
				LastName:  &ln,
			},
			1,
			Expectation{
				Where: "WHERE id = $1 AND first_name = $2 AND last_name = $3",
				Args:  []any{1, "John", "Doe"},
			},
		},
		{
			"mixed pointer fields and non pointer colufieldsmn tags start from 3",
			struct {
				Id        int     `col:"id"`
				FirstName *string `col:"first_name"`
				LastName  *string `col:"last_name"`
			}{
				Id:        1,
				FirstName: &fn,
				LastName:  &ln,
			},
			3,
			Expectation{
				Where: "WHERE id = $3 AND first_name = $4 AND last_name = $5",
				Args:  []any{1, "John", "Doe"},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			where, args := BuildConitionalQuery(tt.arg, tt.startFrom)
			if where != tt.expect.Where {
				t.Errorf("want %s, got %s", tt.expect.Where, where)
			}

			if len(args) != len(tt.expect.Args) {
				t.Errorf("want %v, got %v", tt.expect.Args, args)
			}
		})
	}
}
