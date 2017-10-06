package main

import (
	"testing"
)

func TestSimpleRecursion1e3(t *testing.T) {
	if err := rec(1e3); err != nil {
		t.Errorf("%v", err)
	}
}

func TestSimpleRecursion1e6(t *testing.T) {
	if err := rec(1e6); err != nil {
		t.Errorf("%v", err)
	}
}
