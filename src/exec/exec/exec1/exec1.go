package main

import "fmt"

//1.题目：输入三个整数x,y,z，请把这三个数由小到大输出。
func main() {
	test()

}

func test() {
	var x, y, z, temp float64
	fmt.Println("输入三个数")
	fmt.Scan(&x, &y, &z)
	if x > y {
		temp = x
		x = y
		y = temp
	}
	if x > z {
		temp = z
		z = x
		x = temp
	}
	if y > z {
		temp = y
		y = z
		z = temp
	}
	fmt.Println("从大到小排列", x, y, z)
}
