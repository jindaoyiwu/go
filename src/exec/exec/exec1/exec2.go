package main

import (
	"fmt"
)

//1.题目：输入某年某月某日，判断这一天是这一年的第几天？
//2.程序分析：以3月5日为例，应该先把前两个月的加起来，然后再加上5天即本年的第几天，
func main()  {
	var year,mouth,day int
	fmt.Println("请输入年月日：")
	fmt.Scan(&year, &mouth, &day)

	fmt.Println(test2(year, mouth, day))
}


func test2(y,m,d int) int {
	num := 0
	var monthArr [12]int
	if (y%4 == 0 && y%100 ==0 || y%400 == 0) {
		monthArr = [12]int{31,28,31,30,31,30,31,31,30,31,30,31}
	} else {
		monthArr = [12]int{31,29,31,30,31,30,31,31,30,31,30,31}
	}
	for i:=1; i<=m; i++{
		num += monthArr[i]
	}
	return num + d
}