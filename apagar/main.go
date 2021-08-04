package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

//var localAddr *string = flag.String("l", "0.0.0.0:27016", "local address")
//var remoteAddr *string = flag.String("r", "10.0.0.2:27017", "remote address")

var delayMin, delayMax time.Duration

var inAddress, outAddress string

func main() {
	var err error

	//0.0.0.0:27016
	inAddress = os.Getenv("IN_ADDRESS")

	//10.0.0.2:27017
	outAddress = os.Getenv("OUT_ADDRESS")

	var delayMinAsString = os.Getenv("MIN_DELAY")
	var delayMaxAsString = os.Getenv("MAX_DELAY")
	var delayMinAsInt64 int64
	var delayMaxAsInt64 int64

	delayMinAsInt64, err = strconv.ParseInt(delayMinAsString, 10, 64)
	if err != nil {
		panic(err)
	}

	delayMaxAsInt64, err = strconv.ParseInt(delayMaxAsString, 10, 64)
	if err != nil {
		panic(err)
	}

	delayMin = time.Duration(delayMinAsInt64)
	delayMax = time.Duration(delayMaxAsInt64)

	log.Printf("overloading...")

	proxy()
}

func proxy() {
	flag.Parse()
	fmt.Printf("Listening: %v\nProxying: %v\n\n", inAddress, outAddress)

	listener, err := net.Listen("tcp", inAddress)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		log.Println("New connection", conn.RemoteAddr())
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}
		go func() {
			defer conn.Close()
			conn2, err := net.Dial("tcp", outAddress)
			if err != nil {
				log.Println("error dialing remote addr", err)
				return
			}
			defer conn2.Close()
			closer := make(chan struct{}, 2)
			go copy(closer, conn2, conn)
			go copy(closer, conn, conn2)
			<-closer
			log.Println("Connection complete", conn.RemoteAddr())
		}()
	}
}

func copy(closer chan struct{}, dst io.Writer, src io.Reader) {
	time.Sleep(randDelay(delayMin, delayMax))

	_, _ = io.Copy(dst, src)
	closer <- struct{}{} // connection is closed, send signal to stop proxy
}

func randDelay(min, max time.Duration) (delay time.Duration) {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	delay = max - min
	delay = time.Duration(randGen.Int63n(int64(delay))) + min
	return
}
