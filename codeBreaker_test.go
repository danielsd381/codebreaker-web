package main

import "testing"


func TestCodeBreaker_XXXX(t *testing.T){
	setSecret("1234")
	expected := "XXXX"
	actual:=validate("1234")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
func TestCodeBreakerEmpty(t *testing.T){
	setSecret("1234")
	expected := ""
	actual:=validate("6789")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}

func TestCodeBreaker____(t *testing.T){
	setSecret("1234")
	expected := "----"
	actual:=validate("4321")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
func TestCodeBreaker_(t *testing.T){
	setSecret("1234")
	expected := "-"
	actual:=validate("4567")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
func TestCodeBreaker__(t *testing.T){
	setSecret("1234")
	expected := "--"
	actual:=validate("4367")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
func TestCodeBreaker___(t *testing.T){
	setSecret("1234")
	expected := "---"
	actual:=validate("4327")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
func TestCodeBreakerX_(t *testing.T){
	setSecret("1234")
	expected := "X-"
	actual:=validate("4287")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
func TestCodeBreakerXX__(t *testing.T){
	setSecret("1234")
	expected := "XX--"
	actual:=validate("1324")
	if expected != actual {
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
