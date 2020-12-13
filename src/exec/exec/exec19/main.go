package main

import "fmt"

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	c1 = make(chan int, 1)
	c1 <- 1
	close(c1)
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	var a, b, c chan int
	var j1, j2 int
	a = make(chan int, 1)
	a <-1
	select {
	case j1 =<-c1:
		fmt.Println(11)
	case j2 =<-b:
		fmt.Println('a')
	case j1 =<-c:
		fmt.Println(222)
	default:
	fmt.Println(333)
	}
}