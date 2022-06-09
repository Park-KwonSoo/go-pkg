package str

import (
	"bytes"
)

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
