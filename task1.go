package main

import (
	"fmt"
)

func FizzBuzz(a int) {
	if a%3 == 0 && a%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if a%3 == 0 {
		fmt.Println("Fizz")
	} else if a%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(a)
	}
}
