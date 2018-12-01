package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 1) != 2 {
		t.Error("Expected 1 + 1 to equal 2")
	}
}
