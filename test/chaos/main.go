package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	var counter = 1.0
	ticker := time.NewTicker(500 * time.Millisecond)
	timeout := time.NewTimer(60 * time.Second)
	timeRestart := time.NewTimer(10 * time.Second)

	log.Printf("chaos enable")

	go func() {
		for {
			select {
			case <-timeout.C:
				log.Printf("done!")

			case <-timeRestart.C:
				log.Printf("restart-me!")

			case <-ticker.C:
				counter += 1.0
				log.Printf("counter: %.2f", counter)
			}
		}
	}()

	wg.Wait()

}
