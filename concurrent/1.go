package main

import (
	"fmt"
	"time"
)

//func printRepeatedMessage(message string) {
//	for i := 0; i < 10; i++ {
//		fmt.Println(message)
//	}
//}

func printRepeatedMessage(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
		//time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printRepeatedMessage("Do not communicate by sharing memory; instead, share memory by communicating.")
	go printRepeatedMessage("不要通过共享内存通信，而要通过通信共享内存。")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("END--AuroraSEC")
}
