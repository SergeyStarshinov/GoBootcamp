package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"slices"
	"sync"
	"time"
)

type goroutineInfo struct {
	number int
	time   int
}

func main() {
	goroutinesCount := flag.Int("N", 0, "goroutines count")
	maxTimeToSleep := flag.Int("M", 0, "time to sleep (milliseconds)")
	flag.Parse()
	if *maxTimeToSleep <= 0 || *goroutinesCount <= 0 {
		fmt.Println("two positive parameters were expected: N - goroutines count, M - time to sleep (milliseconds)")
		return
	}
	var wg sync.WaitGroup
	goroutineList := make([]goroutineInfo, 0, *goroutinesCount)

	for i := range *goroutinesCount {
		wg.Add(1)
		timeToSleep := rand.IntN(*maxTimeToSleep)
		goroutineList = append(goroutineList, goroutineInfo{number: i, time: timeToSleep})
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
		}()
	}

	wg.Wait()

	slices.SortFunc(goroutineList, func(a, b goroutineInfo) int {
		return b.time - a.time
	})
	for _, v := range goroutineList {
		fmt.Printf("%d, %d\n", v.number, v.time)
	}

}
