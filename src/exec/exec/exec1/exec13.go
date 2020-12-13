package main

import "fmt"

func main()  {
	d := make(map[string][]int)
	d["a"] = []int{1,2,3}
	d["b"] = []int{4,5,6}
	for _,v := range d{
		v[0] = 3
	}
	fmt.Println(d)
}