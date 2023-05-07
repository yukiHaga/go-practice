package main

import (
	"fmt"
)

func main() {
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)
    ch3 := make(chan int, 1)
		ch1 <- 1
		ch2 <- 2

		select {
		case e1 := <-ch1:
        fmt.Printf("%d: ch1から受信\n", e1)
		case e2 := <-ch2:
        fmt.Printf("%d: ch2から受信\n", e2)
		case ch3 <- 3:
        fmt.Printf("%d: ch3へ送信\n", <-ch3)
		default:
        fmt.Println("ここへは絶対に到達しない")
		}

		a := 1
		b := a

		a = 2

		fmt.Println(a)
		fmt.Println(b)
}