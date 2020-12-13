package main

import "fmt"

//1.题目：求1+2!+3!+...+20!的和
//
//2.程序分析：此程序只是把累加变成了累乘。
//
//3.程序源代码：

func main()  {
	var x int
	fmt.Println("输入一个数")
	fmt.Scan(&x)
	fmt.Println(test9(x))
}

func test9(x int) int {
	a, b := 1,0
	for i:=1; i<=x; i++{
		a  = a*i;
		b += a;
	}
	return b
}