package stream

import "sort"

/*
*	Dev By Kwonsoo
*	Mapper
 */
func Mapper[T any, V any](t T, f func(r T) V) V {
	return f(t)
}

/*
*	Develop By Kwonsoo
*	Map =>  Mapper And Return new Slice
 */
func Map[T any, V any](s []T, f func(T) V) []V {
	r := make([]V, len(s))
	for i, v := range s {
		r[i] = f(v)
	}

	return r
}

/*
*	Develop By Kwonsoo
*	Filter => Return new filterer Slice
 */
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func SortNotChangeOrigin[T any](s []T, f func(i, j int) bool) []T {
	rslt := make([]T, 0)
	for i := 0; i < len(s); i++ {
		rslt = append(rslt, s[i])
	}

	sort.Slice(rslt, f)

	return rslt
}
