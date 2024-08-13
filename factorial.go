package main

import "fmt"

func factorial(n int) int{
	if n == 0 {
		return 1

	}
	return n * factorial(n-1)
}

func main(){
	var num int 
	fmt.Println("Enter the value of num: ")
	fmt.Scan(&num)
	result := factorial(num)
	fmt.Printf("Factorial of %d is %d \n", num , result)
}