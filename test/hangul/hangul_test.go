package hangul_test

import (
	"fmt"
	"testing"

	"github.com/Park-KwonSoo/go-pkg/pkg/hangul"
)

func TestIsChosung(t *testing.T) {
	rslt := hangul.GetCHO("안녕하세요")
	if rslt != "ㅇㄴㅎㅅㅇ" {
		t.Errorf(fmt.Sprintf("It's not string's chosung: %s", rslt))
	}

	rslt = hangul.GetCHO("사과")
	if rslt != "ㅅㄱ" {
		t.Errorf(fmt.Sprintf("It's not string's chosung: %s", rslt))
	}

	rslt = hangul.GetCHO("상자, 갑")
	if rslt != "ㅅㅈ, ㄱ" {
		t.Errorf(fmt.Sprintf("It's not string's chosung: %s", rslt))
	}
}
