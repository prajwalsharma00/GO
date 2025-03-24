package main

import (
	"fmt"
	"strings"
)

func max_key(a map[string]int8) string {
	var max_k string
	var max_value int8 = 0
	for key, value := range a {
		if max_value < value {
			max_value = value
			max_k = key
		}
	}
	return max_k
}
func twomax(a map[string]int8) map[string]int8 {
	var max_k string
	new_map := map[string]int8{}
	max_k = max_key(a)
	new_map[max_k] = a[max_k]
	delete(a, max_k)
	max_k = max_key(a)
	new_map[max_k] = a[max_k]
	delete(a, max_k)
	return new_map

}
func words(words string) map[string]int8 {
	wordscollection := strings.Split(words, " ")
	var wordscount map[string]int8 = map[string]int8{}
	for _, w := range wordscollection {
		wordscount[w]++
	}

	new_map := twomax(wordscount)
	for key, value := range new_map {
		fmt.Printf("%v --> %v\n", key, value)
	}
	return new_map

}
func main() {
	words("hello how how are you hello you how how you  ")

}
