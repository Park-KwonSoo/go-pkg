package stream

import "sort"

/*
*	Develop By Kwonsoo
*	Map => Return new Slice
 */
func Map[T any](s []T, f func(T) T) []T {
	r := make([]T, len(s))
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
