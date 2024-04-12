package util

import (
	"reflect"
	"strconv"
)

func BuildAndConitionalQuery(arg any, startFrom uint) (string, []any) {
	where := ""
	args := []any{}

	rv := reflect.ValueOf(arg)

	for i := 0; i < rv.NumField(); i++ {
		col := rv.Type().Field(i).Tag.Get("col")
		if len(col) == 0 {
			continue
		}

		var v any
		f := rv.Field(i)
		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				continue
			} else {
				v = f.Elem().Interface()
			}
		} else {
			v = rv.Field(i)
		}
		args = append(args, v)
		where += col + " = $" + strconv.Itoa(len(args)+int(startFrom-1)) + " AND "
	}

	if wlen := len(where); wlen > 0 {
		where = "WHERE " + where[:wlen-len(" AND ")] // prepend WHERE and drop the last AND
	}

	return where, args
}
