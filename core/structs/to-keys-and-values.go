package structs

import (
	"reflect"
	"slices"
)

type KeysAndValues struct {
	Keys   []string
	Values []any
}

func ToKeysAndValues(data any, ignoreEmpty bool, ignoreFields *[]string) *KeysAndValues {
	var res KeysAndValues

	v := reflect.ValueOf(data).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		value := v.Field(i).Interface()

		if ignoreFields != nil && slices.Contains(*ignoreFields, tag) {
			continue
		}

		if ignoreEmpty && reflect.DeepEqual(
			value,
			reflect.Zero(v.Field(i).Type()).Interface(),
		) {
			continue
		}

		res.Keys = append(res.Keys, tag)
		res.Values = append(res.Values, value)
	}

	return &res
}
