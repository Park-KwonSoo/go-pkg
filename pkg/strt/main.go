package strt

import (
	"reflect"
)

/*
*	Develop By Kwonsoo
*	TagInfo => Return Struct Tag Info
 */
func GetStructTag[T any](tag string, s *T) []string {
	e := reflect.ValueOf(s).Elem()
	n := e.NumField()

	rslt := make([]string, 0)

	for i := 0; i < n; i++ {
		t := e.Type().Field(i).Tag.Get(tag)
		rslt = append(rslt, t)
	}

	return rslt
}
