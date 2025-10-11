package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Element struct {
	index int
	value int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите первую строку с числами: ")
	if !scanner.Scan() {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	firstSet, err1 := StrToElements(scanner.Text())
	if err1 != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	fmt.Print("Введите вторую строку с числами: ")
	if !scanner.Scan() {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	secondSet, err2 := StrToInts(scanner.Text())
	if err2 != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}

	intersection := FindIntersection(firstSet, secondSet)
	if len(intersection) == 0 {
		fmt.Println("Empty intersection")
	} else {
		fmt.Println(strings.Trim(fmt.Sprint(intersection), "[]"))
	}

}

func FindIntersection(firstSet []Element, secondSet []int) []int {

	slices.SortFunc(firstSet, func(a, b Element) int {
		return a.value - b.value
	})
	slices.Sort(secondSet)

	intersection := []Element{}
	i1 := 0
	i2 := 0
	for i1 < len(firstSet) && i2 < len(secondSet) {
		if firstSet[i1].value == secondSet[i2] {
			intersection = append(intersection, firstSet[i1])
			i1++
			i2++
		} else if firstSet[i1].value < secondSet[i2] {
			i1++
		} else {
			i2++
		}
	}
	// сортировка по первоначальным индексам
	slices.SortFunc(intersection, func(a, b Element) int {
		return a.index - b.index
	})

	result := make([]int, len(intersection))
	for i, el := range intersection {
		result[i] = el.value
	}
	return result
}

func StrToElements(text string) ([]Element, error) {
	tokens := strings.Fields(text)
	result := []Element{}
	for i, token := range tokens {

		num, err := strconv.Atoi(token)
		if err != nil {
			return []Element{}, errors.New(token)
		}
		result = append(result, Element{index: i, value: num})
	}
	return result, nil
}

func StrToInts(text string) ([]int, error) {
	tokens := strings.Fields(text)
	result := []int{}
	for _, token := range tokens {

		num, err := strconv.Atoi(token)
		if err != nil {
			return []int{}, errors.New(token)
		}
		result = append(result, num)
	}
	return result, nil
}
