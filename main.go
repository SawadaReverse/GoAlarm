package main

import (
	"fmt"
	"time"
)

func main() {
	var SetTime string
	var nowTime string
	fmt.Println("Enter time (HH:mm:ss)")
	fmt.Scanf("%s", &SetTime)
	fmt.Println(SetTime)

	for {
		nowTime = time.Now().Format("15:04:05")
		if nowTime == SetTime {
			fmt.Println("It's time !!!")
			return
		}

	}
}
