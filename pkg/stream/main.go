package stream

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
