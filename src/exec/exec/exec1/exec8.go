package main

import "fmt"

//利用条件运算符的嵌套来完成此题：
//学习成绩>=90分的同学用A表示，60-89分之间的用B表示，60分以下的用C表示。
func main(){
	var score float32
	fmt.Println("请写出你的成绩")
	fmt.Scan(&score)
	if score <0 || score >100 {
		fmt.Println("非法")
	}
	if score > 90{
		fmt.Println("你的成绩是A")
	}
	if score >= 90 && score<60 {
		fmt.Println("你的成绩是B")
	}
	if score <=60 {
		fmt.Println("你的成绩是C")
	}
}
