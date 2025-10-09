package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	text, errText := inputText()
	if errText != nil {
		log.Fatalf("runtime error: %s", errText)
	}
	k, errInt := inputNumber()
	if errInt != nil {
		log.Fatalf("runtime error: %s", errInt)
	}

	fmt.Println(PrintWords(text, k))

}

func PrintWords(text string, k int) string {
	countMap := invertMap(countWords(text))

	keys := sortMapKeysDesc(countMap)

	result := []string{}
	printed := 0
	for _, key := range keys {
		slice := countMap[key]
		slices.Sort(slice)
		for _, str := range slice {
			if printed >= k {
				return strings.Join(result, " ")
			}
			result = append(result, str)
			printed++
		}
	}

	return strings.Join(result, " ")
}

func sortMapKeysDesc(countMap map[int][]string) []int {
	keys := make([]int, len(countMap))
	i := 0
	for key := range countMap {
		keys[i] = key
		i++
	}
	slices.SortFunc(keys, func(a, b int) int {
		return b - a
	})
	return keys
}

func countWords(text string) map[string]int {
	slice := strings.Split(text, " ")
	words := map[string]int{}
	for _, str := range slice {
		words[str]++
	}
	return words
}

func invertMap(sourceMap map[string]int) map[int][]string {
	result := map[int][]string{}
	for key, value := range sourceMap {
		result[value] = append(result[value], key)
	}
	return result
}

func inputText() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	if !scanner.Scan() {
		return "", fmt.Errorf("error reading line")
	}
	return scanner.Text(), nil
}

func inputNumber() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число K, ограничение на количество слов: ")
	if !scanner.Scan() {
		return 0, fmt.Errorf("error reading line")
	}
	result, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("it is not a number")
	}
	return result, nil
}
