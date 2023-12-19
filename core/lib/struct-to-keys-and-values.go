package lib

import (
	"reflect"
)

type response struct {
	Keys   []string
	Values []any
}

func StructToKeysAndValues(data any, ignoreEmpty, ignoreID bool) *response {
	var res response

	v := reflect.ValueOf(data).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		value := v.Field(i).Interface()

		if tag == "id" && ignoreID {
			continue
		}

		if ignoreEmpty && reflect.DeepEqual(value, reflect.Zero(v.Field(i).Type()).Interface()) {
			continue
		}

		res.Keys = append(res.Keys, tag)
		res.Values = append(res.Values, value)
	}

	return &res
}
