package main

import "fmt"

func main() {
	var arr [3]int8 = [3]int8{1, 2, 3}
	for _, i := range arr {
		if i%2 == 0 {
			fmt.Println("this is even  %v", i)
		} else {
			fmt.Println("this is odd %v", i)
		}
	}
	var code map[string]uint8 = map[string]uint8{"hello": 8}
	var age, ok = code["hi"]
	if ok == false {
		fmt.Printf("the value is worng for taht %v ", age)
	}

}
