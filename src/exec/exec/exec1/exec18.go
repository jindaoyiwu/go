package main


import (
	"fmt"
)

//func main() {
//	a := []int{1, 2, 3}
//	for _, i := range a {
//		fmt.Println(i)
//			defer p(i)
//		}
//}
//
//func p(i int) {
//	fmt.Println(i)
//}

func main() {
	a := []int{1, 2, 3}
	for _, i := range a {
		fmt.Println(i)
		defer func() {
			fmt.Println(i)
		}()
	}
}

