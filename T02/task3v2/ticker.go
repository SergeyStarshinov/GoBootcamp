package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())

	go ticker(*tickerInterval, ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	cancel()
	fmt.Println("Termination")

}

func ticker(interval uint, ctx context.Context) {
	var timeFromStart uint = 0
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("Tick %d since %d\n", i, timeFromStart)
			time.Sleep(time.Duration(interval) * time.Second)
			timeFromStart += interval
		}
	}
}
