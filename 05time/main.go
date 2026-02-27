package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006"))
	fmt.Println(currentTime.Format("01-02-2006 Monday"))
	fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:05")) // for format the pattern is fixed "01-02-2006 Monday 15:04:05"

	createdTime := time.Date(2026, time.February, 26, 18, 8, 0, 0, time.UTC)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("01-02-03 Monday"))
}
