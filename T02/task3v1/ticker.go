package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	tickerInterval := flag.Uint("K", 0, "ticker interval in seconds")
	flag.Parse()
	if *tickerInterval == 0 {
		fmt.Println("the ticker interval can't be zero")
		return
	}

	stopCh := make(chan int)
	go ticker(*tickerInterval, stopCh)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	stopCh <- 1
	fmt.Println("Termination")

}

func ticker(interval uint, stopCh <-chan int) {
	var timeFromStart uint = 0
	for i := 0; ; i++ {
		select {
		case <-stopCh:
			return
		default:
			fmt.Printf("Tick %d since %d\n", i, timeFromStart)
			time.Sleep(time.Duration(interval) * time.Second)
			timeFromStart += interval
		}
	}
}
