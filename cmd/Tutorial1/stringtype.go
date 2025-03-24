package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string = "hi "
	fmt.Println(utf8.RuneCountInString(name))
}
