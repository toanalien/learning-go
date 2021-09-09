package main

import "fmt"

func findMaxLength(str []string) (max int) {
	for _, s := range str {
		if len(s) > max {
			max = len(s)
		}
	}
	return
}

func findStringWithLength(str []string, length int) (returnArr []string) {
	for _, s := range str {
		if len(s) == length {
			returnArr = append(returnArr, s)
		}
	}
	return
}

func main() {
	var str = []string{"aba", "aa", "ad", "c", "vcd"}
	maxLength := findMaxLength(str)
	fmt.Println(findStringWithLength(str, maxLength))
}
