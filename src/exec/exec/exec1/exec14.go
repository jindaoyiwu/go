package main

import "fmt"

func main()  {
	var arr1 []*int
	arr1 = make([]*int, 0)

	arr2 := make(map[string]int)
	arr2["a"] = 1
	arr2["b"] = 2
	for _, v := range arr2{
		arr1 = append(arr1, &v)
	}

	arr3 := make([]int, len(arr2))
	for i,_ := range arr3{
		fmt.Println(*arr1[i])
		arr3[i] = *arr1[i]
	}
	fmt.Println(arr3)
	fmt.Println(arr2)
}