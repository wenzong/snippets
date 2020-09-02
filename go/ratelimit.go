package main

import (
	"fmt"
	"time"
)

func main() {

	rate := time.Second * 10
	throttle := time.Tick(rate)
	for {
		<-throttle // rate limit our Service.Method RPCs
		fmt.Println(time.Now())
	}
}
