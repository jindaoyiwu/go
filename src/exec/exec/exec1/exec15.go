package main

import "fmt"

//阶乘
func main()  {
	fmt.Println(factorial(3))
}
func factorial(n int) int {
	if (n>1) {
		return n*factorial(n-1)
	}
	return 1
}

// 费波南 1,1,2,3,5,8,13,21
func fbn(n int) int {
	if n ==1 || n == 2{
		return 1
	} else {
		return fbn(n-1)+fbn(n-2)
	}
}