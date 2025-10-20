package main

import (
	"flag"
	"fmt"
)

func main() {
	startRange := flag.Int("K", 0, "start of the range to generate (0 by default)")
	endRange := flag.Int("N", 0, "end of the range to generate (0 by default)")
	flag.Parse()
	if *startRange > *endRange {
		fmt.Printf("the start (K=%d) can't be less than the end (N=%d) of the range\n", *startRange, *endRange)
		return
	}

	ch1 := generate(*startRange, *endRange)
	ch2 := square(ch1)

	for val := range ch2 {
		fmt.Println(val)
	}

}

func generate(startRange, endRange int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := startRange; i <= endRange; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func square(inputCh <-chan int) <-chan int {
	outputCh := make(chan int)
	go func(inputCh <-chan int, outputCh chan<- int) {
		for val := range inputCh {
			squared := val * val
			outputCh <- squared
		}
		close(outputCh)
	}(inputCh, outputCh)
	return outputCh
}
