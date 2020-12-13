package main

import "fmt"

//1.题目：求s=a+aa+aaa+aaaa+aa...a的值，其中a是一个数字。例如2+22+222+2222+22222(此时共有5个数相加)，几个数相加有键盘控制。
func main()  {
	var n,m,sn= 1, 1, 0
	fmt.Println("请输入一个数字：")
	fmt.Scan(&n)
	fmt.Println("请输入循环数字：")
	fmt.Scan(&m)

	for j:=1; j<=m; j++{
		fmt.Println(num(n, j))
		sn += num(n, j)
	}

	fmt.Println(sn)
}

func num(n,m int) int {
	kn := n
	for i:=0; i<m; i++{
		kn += n*exp10(i)
	}
	return kn
}

func exp10( n int) int {
	if n == 0 {
		return 0
	}
	m := 1
	for i:=0; i<n; i++{
		m *= 10
	}
	return m
}