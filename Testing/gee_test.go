package main

import (
	"testing"
)

// test function
func TestReturnHello(t *testing.T) {
	actualString := ReturnHello()
	expectedString := "Hello"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
