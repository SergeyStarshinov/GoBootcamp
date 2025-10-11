package main

import (
	"slices"
	"testing"
)

func TestStrToElementsEmpty(t *testing.T) {
	text := ""
	result, err := StrToElements(text)
	if err != nil {
		t.Errorf("Result was incorrect, got error: %s, want: nil.", err)
	} else if len(result) != 0 {
		t.Errorf("Result was incorrect, got slice with %d elements, want: empty slice.", len(result))
	}
}

func TestStrToElementsError(t *testing.T) {
	text := "1 2 a 4"
	_, err := StrToElements(text)
	if err == nil {
		t.Errorf("Result was incorrect, got nil, want: error.")
	}
}

func TestStrToElements(t *testing.T) {
	text := "1 -1 0"
	expected := []Element{
		{index: 0, value: 1},
		{index: 1, value: -1},
		{index: 2, value: 0},
	}
	result, err := StrToElements(text)
	if err != nil {
		t.Errorf("Result was incorrect, got error: %s, want: nil.", err)
	}
	if !slices.EqualFunc(result, expected, func(a, b Element) bool {
		if a.index == b.index && a.value == b.value {
			return true
		}
		return false
	}) {
		t.Errorf("Result was incorrect, got: %+v, want: %+v.", result, expected)
	}
}

func TestStrToElementsSpaces(t *testing.T) {
	text := "1  -1  0"
	expected := []Element{
		{index: 0, value: 1},
		{index: 1, value: -1},
		{index: 2, value: 0},
	}
	result, err := StrToElements(text)
	if err != nil {
		t.Errorf("Result was incorrect, got error: %s, want: nil.", err)
	}
	if !slices.EqualFunc(result, expected, func(a, b Element) bool {
		if a.index == b.index && a.value == b.value {
			return true
		}
		return false
	}) {
		t.Errorf("Result was incorrect, got: %+v, want: %+v.", result, expected)
	}
}

func TestStrToIntsEmpty(t *testing.T) {
	text := ""
	result, err := StrToInts(text)
	if err != nil {
		t.Errorf("Result was incorrect, got error: %s, want: nil.", err)
	} else if len(result) != 0 {
		t.Errorf("Result was incorrect, got slice with %d elements, want: empty slice.", len(result))
	}
}

func TestStrToIntsError(t *testing.T) {
	text := "1 2 a 4"
	_, err := StrToInts(text)
	if err == nil {
		t.Errorf("Result was incorrect, got nil, want: error.")
	}
}

func TestStrToInts(t *testing.T) {
	text := "1 -1 0"
	expected := []int{1, -1, 0}
	result, err := StrToInts(text)
	if err != nil {
		t.Errorf("Result was incorrect, got error: %s, want: nil.", err)
	}
	if !slices.Equal(result, expected) {
		t.Errorf("Result was incorrect, got: %+v, want: %+v.", result, expected)
	}
}

func TestStrToIntsSpaces(t *testing.T) {
	text := "1  -1  0"
	expected := []int{1, -1, 0}
	result, err := StrToInts(text)
	if err != nil {
		t.Errorf("Result was incorrect, got error: %s, want: nil.", err)
	}
	if !slices.Equal(result, expected) {
		t.Errorf("Result was incorrect, got: %+v, want: %+v.", result, expected)
	}
}

func TestFindIntersectionEmpty(t *testing.T) {
	text1 := "1 2"
	text2 := "3 4"
	firstSet, _ := StrToElements(text1)
	seconSet, _ := StrToInts(text2)
	intersection := FindIntersection(firstSet, seconSet)
	if len(intersection) != 0 {
		t.Errorf("Result was incorrect, got intersection slice with %d elements, want: empty slice.",
			len(intersection))
	}
}

func TestFindIntersectionEmptyText1(t *testing.T) {
	text1 := ""
	text2 := "3 4"
	firstSet, _ := StrToElements(text1)
	seconSet, _ := StrToInts(text2)
	intersection := FindIntersection(firstSet, seconSet)
	if len(intersection) != 0 {
		t.Errorf("Result was incorrect, got intersection slice with %d elements, want: empty slice.",
			len(intersection))
	}
}

func TestFindIntersectionEmptyText2(t *testing.T) {
	text1 := "1"
	text2 := " "
	firstSet, _ := StrToElements(text1)
	seconSet, _ := StrToInts(text2)
	intersection := FindIntersection(firstSet, seconSet)
	if len(intersection) != 0 {
		t.Errorf("Result was incorrect, got intersection slice with %d elements, want: empty slice.",
			len(intersection))
	}
}

func TestFindIntersectionNotEmpty(t *testing.T) {
	text1 := "5 3 4 2 1 6"
	text2 := "6 4 2 4"
	expected := []int{4, 2, 6}
	firstSet, _ := StrToElements(text1)
	seconSet, _ := StrToInts(text2)
	intersection := FindIntersection(firstSet, seconSet)
	if !slices.Equal(intersection, expected) {
		t.Errorf("Result was incorrect, got: %+v, want: %+v.", intersection, expected)
	}
}
