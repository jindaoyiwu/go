package main

import "fmt"

//1.判断101-200之间有多少个素数，并输出所有素数。
//
//2.程序分析：判断素数的方法：用一个数分别去除2，如果能被整除，则表明此数不是素数，反之是素数。
func main()  {
	I:
	for n := 2; n <= 10; n++ {
		for m := 2; m<n; m++{
			if n%m == 0{
				continue I
			}
		}
		fmt.Println(n)
	}

}