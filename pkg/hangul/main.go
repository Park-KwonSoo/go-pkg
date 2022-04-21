package hangul

import "bytes"

var (
	hangulCHO = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
	hangulJUN = []string{"ㅏ", "ㅐ", "ㅑ", "ㅒ", "ㅓ", "ㅔ", "ㅕ", "ㅖ", "ㅗ", "ㅘ", "ㅙ", "ㅚ", "ㅛ", "ㅜ", "ㅝ", "ㅞ", "ㅟ", "ㅠ", "ㅡ", "ㅢ", "ㅣ"}
	hangulJON = []string{"", "ㄱ", "ㄲ", "ㄳ", "ㄴ", "ㄵ", "ㄶ", "ㄷ", "ㄹ", "ㄺ", "ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ", "ㅀ", "ㅁ", "ㅂ", "ㅄ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
)

const (
	hangulBASE = rune('가') //첫번째 글자
	hangulEND  = rune('힣') //마지막 글자

	hangulJA = rune('ㄱ') //자음
	hangulMO = rune('ㅏ') //모음
)

/*
*	Dev By Kyle
*	한글 문자열을 초성 문자열로 변환시킨다.
 */
func GetCHO(h string) string {

	//한글문자열의 길이
	var b bytes.Buffer
	for _, char := range []rune(h) {
		if char >= hangulBASE && char <= hangulEND {
			temp := char - hangulBASE
			b.WriteString(hangulCHO[temp/588])
		}
	}

	return b.String()
}
