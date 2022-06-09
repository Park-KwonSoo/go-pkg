package str_test

import (
	"testing"

	"github.com/Park-KwonSoo/go-pkg/pkg/str"
)

func TestConcat(t *testing.T) {
	rslt := str.Concat("Hello ", "World!")
	if rslt != "Hello World!" {
		t.Errorf("It's not concated")
	}
}
