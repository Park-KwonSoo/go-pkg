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

func TestOf(t *testing.T) {

	rslt := str.Of("hello", 3)
	if rslt != "l" {
		t.Errorf("invalid func")
	}
}
