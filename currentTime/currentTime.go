package main

import (
	"fmt"
	"time"
)

func main() {

	timeMilliSeconds := time.Now().UnixNano() / int64(time.Millisecond)
	totalSeconds := timeMilliSeconds / 1000
	currentSeconds := totalSeconds % 60
	totalMinutes := totalSeconds / 60
	currentMinutes := totalMinutes % 60
	totalHours := totalMinutes / 60
	currentHours := totalHours % 60

	fmt.Printf("TotalMilliSeconds: %d\n", timeMilliSeconds)
	fmt.Printf("TotalSeconds: %d\n", totalSeconds)
	fmt.Printf("CurrentSeconds: %d\n", currentSeconds)
	fmt.Printf("TotalMinutes: %d\n", totalMinutes)
	fmt.Printf("CurrentMinutes: %d\n", currentMinutes)
	fmt.Printf("TotalHours: %d\n", totalHours)
	fmt.Printf("CurrentHours: %d\n", currentHours)

}
