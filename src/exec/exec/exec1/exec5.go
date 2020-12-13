package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%c%c", 219, 219)
		}
		fmt.Println()
	}
}