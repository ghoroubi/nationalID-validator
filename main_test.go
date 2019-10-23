package main

import "testing"

func TestNationalCodeIsValid(t *testing.T) {
	code1 := "2949850685"
	code2 := "2949850686"
	if !NationalCodeIsValid(code1) {
		t.FailNow()
	}
	if NationalCodeIsValid(code2) {
		t.FailNow()
	}
}
