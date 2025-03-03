package schema

import (
	"reflect"
)

// return pointers slices for struct fields
func GetFieldPointers[T any](s *T) []any {
	v := reflect.ValueOf(s).Elem()
	numFields := v.NumField()
	pointers := make([]any, numFields)

	for i := range numFields {
		pointers[i] = v.Field(i).Addr().Interface()
	}

	return pointers
}
