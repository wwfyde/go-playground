package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start count")
	tick := time.Tick(1 * time.Second)
	for range time.Tick(1 * time.Second) {
		print("hello!")
	}
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	fmt.Println("countdown")

}
