package main

import "fmt"

func factorial(a int) int {
	if a == 0 {
		return 1
	} else {
		return a * factorial(a-1)
	}
}

func main() {

	var fact int = factorial(5)
	fmt.Printf("the factorial of 5 is %d\n", fact)

}
