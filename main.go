package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("Start time is", startTime)

	finishTime := startTime.Add(time.Second * 5)
	fmt.Println("Finish time is", finishTime)

	currTime := startTime

	for i := 0; currTime.Before(finishTime); i++ {
		currTime = time.Now()
		fmt.Println("Current time is:", currTime)

	}

	fmt.Println("Finished")
}