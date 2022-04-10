package validator_test

import (
	"testing"

	"github.com/Park-KwonSoo/go-pkg/pkg/validator"
)

func TestIsEmail(t *testing.T) {
	rslt := validator.IsEmail("ajsdlk@naver.com")
	if !rslt {
		t.Errorf("It's Email Right")
	}

	rslt = validator.IsEmail("alsdjl@naver")
	if rslt {
		t.Errorf("1: It's Not Email Right")
	}

	rslt = validator.IsEmail("alsdjl")
	if rslt {
		t.Errorf("2: It's Not Email Right")
	}

	rslt = validator.IsEmail("naver.com")
	if rslt {
		t.Errorf("3: It's Not Email Right")
	}

	rslt = validator.IsEmail("alsdjl@.com")
	if rslt {
		t.Errorf("4: It's Not Email Right")
	}

	rslt = validator.IsEmail("@naver.com")
	if rslt {
		t.Errorf("5: It's Not Email Right")
	}

	rslt = validator.IsEmail("@.")
	if rslt {
		t.Errorf("6: It's Not Email Right")
	}
}

func TestIsRightPassword(t *testing.T) {
	rslt := validator.IsRightPassword("asdjasd", 8)
	if rslt {
		t.Errorf("1. It's not right password")
	}
}
