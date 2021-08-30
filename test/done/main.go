package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	timeout := time.NewTimer(10 * time.Second)

	go func() {
		for {
			select {
			case <-timeout.C:
				log.Printf("done!")
			}
		}
	}()

	wg.Wait()

}
