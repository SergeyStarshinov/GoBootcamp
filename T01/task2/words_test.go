package main

import (
	"testing"
)

func TestEmptyString(t *testing.T) {
	text := ""
	k := 3
	expected := ""
	result := PrintWords(text, k)
	if expected != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestZeroK(t *testing.T) {
	text := "a b c a"
	k := 0
	expected := ""
	result := PrintWords(text, k)
	if expected != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestSmallK(t *testing.T) {
	text := "aa bb cc aa cc cc cc aa ab ac bb"
	k := 3
	expected := "cc aa bb"
	result := PrintWords(text, k)
	if expected != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestBigK(t *testing.T) {
	text := "aa bb aa"
	k := 5
	expected := "aa bb"
	result := PrintWords(text, k)
	if expected != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestK1(t *testing.T) {
	text := "aa aa bb bb"
	k := 1
	expected := "aa"
	result := PrintWords(text, k)
	if expected != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}
