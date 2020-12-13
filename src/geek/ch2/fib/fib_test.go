package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var a int = 1
	var b int = 1
	fmt.Println(a)
	for i:=1; i<5; i++{
		fmt.Println(b)
		tmp := a
		a = b
		b = tmp + a

	}
	fmt.Println()
}
