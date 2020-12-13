package main

import (
	"fmt"
)

//1.打印出所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方和等于该数本身。
//todo 优化思路：当不可能两个数字同时大于8
func main()  {
	test7()
}

func test7()  {
	for i:=1; i<10; i++{
		for j:=1; j<10; j++{
			for k:=1; k<10; k++{
				sum := i*100+j*10+k;
				if (sum == i*i*i + j*j*j + k*k*k) {
					fmt.Printf("水仙花数 %d", sum)
					fmt.Println()
				}
			}
		}
	}
}
