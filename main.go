package main

import "fmt"

func main() {
	// task1
	fmt.Println("---------------TASK1--------------------")
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
	//task2
	fmt.Println("---------------TASK2--------------------")
	dobStr := "28.07.1997" // Replace this date with your birthday
	task2(dobStr)

	//task3
	fmt.Println("---------------TASK3--------------------")
	DisplayNumberInReverseOrderWithDefer()

	//task5
	fmt.Println("---------------TASK5--------------------")
	var a int
	fmt.Println("Enter any integer number: ")
	fmt.Scan(&a)
	fmt.Println(recur(a))
}
