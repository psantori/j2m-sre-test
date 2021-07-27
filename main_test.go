package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Hello, visitor number 42!"
	actual := greet(42)
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
