package main

import "fmt"

//1,1,2,3,5,8,13,21 计算前n个数的和
func main()  {
	var num int
	fmt.Println("一个数字：")
	fmt.Scan(&num)

	println(test6(num), sum)
}
var sum =0
func test6(n int) int {
	if n ==1 {
		sum = 1
		return 1
	}
	if n ==2 {
		sum = 2
		return 1
	}
	t := test6(n-2) + test6(n-1)
	sum = sum + t
	return t
}