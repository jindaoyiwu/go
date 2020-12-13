package main

import "fmt"

func main()  {
	var i int
	fmt.Println("输入一个数字")
	fmt.Scan(&i)
	test11(i)
}

func test11(j int)  {
	t := 100
	sum := 100
	for i:=1; i<j; i++{
		t = t/2
		sum += t*2
	}
	fmt.Println(sum)
}