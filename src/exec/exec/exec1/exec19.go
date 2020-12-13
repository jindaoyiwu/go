package main

import "fmt"

func main()  {
	var a []int
	a = []int{3,4,2,1,7,4,12,20}
	bubbleSort(a)
}

func bubbleSort(a []int)  {
	len := len(a)
	if len <= 1{
		fmt.Println(a)
		return
	}
	for i:=0; i<len; i++{
		for j:=0; j<len-i-1; j++{
			if (a[j] > a[j+1]) {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}