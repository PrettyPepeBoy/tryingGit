package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func GetRequest() *http.Request {
	return nil
}

func SendResponse(*http.Request, *http.Response) {

}

func DoSomeStuff(*http.Request) *http.Response {
	return nil
}

func ServeRequest(routineNumber int, done chan struct{}) {
serveLoop:
	for {
		select {
		case <-done:
			break serveLoop

		default:
		}

		req := GetRequest()
		resp := DoSomeStuff(req)
		SendResponse(req, resp)

		time.Sleep(time.Duration(rand.Uint32()%100) * time.Millisecond * 10)
		log.Println(routineNumber)
	}

	log.Println("Done serving", routineNumber)
}

func Sender(c chan int, done chan struct{}) {
	counter := 0
	for {
		select {
		case c <- counter:
			counter++

		case <-done:
			close(c)
			fmt.Println("Sender done")
			return
		}
	}
}

func Receiver(c chan int) {
	for i := range c {
		fmt.Println(i * i)
	}
	fmt.Println("Receiver done")
}

func main() {
	c := make(chan int)
	done := make(chan struct{})

	go Sender(c, done)
	go Receiver(c)

	time.Sleep(time.Second * 3)
	close(done)

	sigC := make(chan os.Signal)
	signal.Notify(sigC, os.Interrupt)

	sig := <-sigC
	fmt.Println(sig.String())
}
