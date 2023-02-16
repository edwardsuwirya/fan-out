package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chProducer := make(chan int)
	chConsumer1 := make(chan int)
	chConsumer2 := make(chan int)

	go producer(chProducer)
	go consumer(chConsumer1, "A")
	go consumer(chConsumer2, "B")
	fanOut(chProducer, chConsumer1, chConsumer2)
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}

func fanOut(chProducer chan int, chConsumer1, chConsumer2 chan<- int) {
	for n := range chProducer {
		if n < 50 {
			chConsumer1 <- n
			//close(chProducer)
		} else {
			chConsumer2 <- n
		}
	}
}

func producer(ch chan<- int) {
	for {
		sleep()
		n := rand.Intn(100)
		fmt.Printf(" -> %d\n", n)
		ch <- n
	}
}
func consumer(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("consumer %s <- %d\n", name, n)
	}
}
