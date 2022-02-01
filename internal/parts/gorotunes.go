package parts

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func Gorutunes() {

	var c chan string = make(chan string)

	go ponger(c)
	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
	