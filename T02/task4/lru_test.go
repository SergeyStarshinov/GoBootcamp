package main

import (
	"sync"
	"testing"

	"t02/task4/lrucache"
)

func TestEmptyGet(t *testing.T) {
	myCache := lrucache.NewMyCache[int](2)
	expected := 0
	result, exists := myCache.Get("a")
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
	if exists {
		t.Errorf("Result was incorrect, got: %d, expected: not exists.", result)
	}
}

func TestSet(t *testing.T) {
	myCache := lrucache.NewMyCache[string](2)
	myCache.Set("a", "aaa")
	myCache.Set("a", "bbb")
	result, exists := myCache.Get("a")
	expected := "aaa"
	if expected != result {
		t.Errorf("Result was incorrect, got: %s, expected: %s.", result, expected)
	}
	if !exists {
		t.Errorf("Result was incorrect, got: not exists, expected: %s.", expected)
	}
}

func TestSetSlices(t *testing.T) {
	myCache := lrucache.NewMyCache[[]int](2)
	myCache.Set("a", []int{1})
	myCache.Set("b", []int{2, 3})
	result, _ := myCache.Get("a")
	if result == nil {
		t.Errorf("Result was incorrect, got: nil, expected: %v.", result)
	}
}

func TestSetWithRemove(t *testing.T) {
	myCache := lrucache.NewMyCache[int](2)
	myCache.Set("1", 1)
	myCache.Set("2", 2)

	result, _ := myCache.Get("1")
	expected := 1
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}

	myCache.Set("3", 3)

	result, _ = myCache.Get("2")
	expected = 0
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
	result, _ = myCache.Get("3")
	expected = 3
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
	result, _ = myCache.Get("1")
	expected = 1
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
}

func TestGoroutines(t *testing.T) {
	myCache := lrucache.NewMyCache[int](2)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		myCache.Get("0")
		wg.Done()
	}()
	go func() {
		myCache.Set("1", 1)
		wg.Done()
	}()
	go func() {
		myCache.Set("2", 2)
		wg.Done()
	}()
	wg.Wait()

	expected := 1
	result, _ := myCache.Get("1")
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
	expected = 2
	result, _ = myCache.Get("2")
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}

	wg.Add(2)
	go func() {
		myCache.Set("3", 3)
		wg.Done()
	}()
	go func() {
		myCache.Set("4", 4)
		wg.Done()
	}()
	wg.Wait()

	expected = 0
	result, _ = myCache.Get("1")
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
	expected = 0
	result, _ = myCache.Get("2")
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
	expected = 3
	result, _ = myCache.Get("3")
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
	expected = 4
	result, _ = myCache.Get("4")
	if expected != result {
		t.Errorf("Result was incorrect, got: %d, expected: %d.", result, expected)
	}
}
