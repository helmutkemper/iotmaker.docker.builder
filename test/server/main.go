package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	//fixme: apagar - início
	var counter = 1.0
	var memory []byte = make([]byte, 0)
	tk := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-tk.C:
				counter += 1.321
				log.Printf("blablabla counter: %.2f", counter)
				memory = append(memory, make([]byte, 500*1024*1024)...)
			}
		}
	}()
	//fixme: apagar - fim

	fmt.Printf("starting server at port 3000\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
