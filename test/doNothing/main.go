package main

import (
	"log"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	log.Printf("done!")

	go func() {
		time.Sleep(60 * time.Second)
		wg.Done()
	}()

	wg.Add(1)
	wg.Wait()
}
