package main

import (
	"testing"
)

func TestMyPow(t *testing.T) {
	res := myPow(4, 3)
	expected := 64.0
	if res != expected {
		t.Fatalf("expected %f, got %f", res, expected)
	}
}
