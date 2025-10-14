package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Visit struct {
	spec string
	date string
}

type Storage map[string][]Visit

type UserNotFoundError struct {
	message string
}

func (e UserNotFoundError) Error() string {
	return e.message
}

func newError(message string) *UserNotFoundError {
	return &UserNotFoundError{message: message}
}

func main() {
	storage := Storage{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command (Save, GetHistory, GetLastVisit or Exit): ")
		if !scanner.Scan() {
			log.Fatalf("invalid input")
		}
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))
		var err error
		switch command {
		case "save":
			err = storage.SaveData()
		case "gethistory":
			err = storage.GetHistory()
		case "getlastvisit":
			err = storage.GetLastVisit()
		case "exit":
			return
		default:
			fmt.Println("unknown command, try again")
		}
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}
}

func (s Storage) SaveData() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Visitor's full name: ")
	if !scanner.Scan() {
		return fmt.Errorf("name input error")
	}
	visitorName := strings.TrimSpace(scanner.Text())

	fmt.Print("Doctor's specialization: ")
	if !scanner.Scan() {
		return fmt.Errorf("specialization input error")
	}
	specialization := strings.ToLower(strings.TrimSpace(scanner.Text()))

	fmt.Print("Visit date (YYYY-MM-DD): ")
	if !scanner.Scan() {
		return fmt.Errorf("date input error")
	}
	layout := "2006-01-02"
	if _, err := time.Parse(layout, scanner.Text()); err != nil {
		return fmt.Errorf("incorrect date format")
	}
	s[visitorName] = append(s[visitorName], Visit{spec: specialization, date: scanner.Text()})
	return nil
}

func (s Storage) GetHistory() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Visitor's full name: ")
	if !scanner.Scan() {
		return fmt.Errorf("name input error")
	}
	visitorName := strings.TrimSpace(scanner.Text())
	visits, ok := s[visitorName]
	if !ok {
		return newError("patient not found")
	}
	for _, visit := range visits {
		fmt.Printf("%s %s\n", visit.spec, visit.date)
	}
	return nil
}

func (s Storage) GetLastVisit() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Visitor's full name: ")
	if !scanner.Scan() {
		return fmt.Errorf("name input error")
	}
	visitorName := strings.TrimSpace(scanner.Text())
	visits, ok := s[visitorName]
	if !ok {
		return newError("patient not found")
	}

	fmt.Print("Doctor's specialization: ")
	if !scanner.Scan() {
		return fmt.Errorf("specialization input error")
	}
	specialization := strings.ToLower(strings.TrimSpace(scanner.Text()))

	visitSpec := []Visit{}
	for _, visit := range visits {
		if visit.spec == specialization {
			visitSpec = append(visitSpec, visit)
		}
	}
	if len(visitSpec) == 0 {
		fmt.Printf("%s didn't visit %s\n", visitorName, specialization)
	} else {
		slices.SortFunc(visitSpec, func(a, b Visit) int {
			if a.date > b.date {
				return -1
			} else if a.date < b.date {
				return 1
			}
			return 0
		})
		fmt.Println(visitSpec[0].date)
	}
	return nil
}
