package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string = "hi "
	var myBoolean bool = true
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Println(myBoolean)

}
