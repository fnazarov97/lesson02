package main

import "fmt"

func DisplayNumberInReverseOrderWithDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
	}
}
