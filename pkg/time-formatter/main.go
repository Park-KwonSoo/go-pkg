package time_parshing

import (
	"time"
)

/*
*	Dev By Kyle
*	time을 YYYY-MM-DD HH:mm:ss로 변환
 */
func Parse_YYMMDD_HHMMSS(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

/*
*	Dev By Kyle
*	time을 YYYY-MM-DD로 변환
 */
func Parse_YYMMDD(t time.Time) string {
	return t.Format("2006-01-02")
}
