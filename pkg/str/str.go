package str

import (
	"bytes"
)

type S struct {
	string
}

/*
*	Dev By Kwonsoo
*	문자열 합치는 함수
 */
func Concat(s ...string) string {
	var b bytes.Buffer

	for i := 0; i < len(s); i++ {
		b.WriteString(s[i])
	}

	return b.String()
}

/*
*	Dev By Kwonsoo
*	문자열의 인덱스 값을 가져오는 함수
 */
func Of(s string, index int) string {
	return string(s[index])
}
