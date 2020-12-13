package main

import "time"

func main() {
	//test17_1()
	//test17_2()
	test17_3()
}

func test17_1() {
	println("start")
	ch := make(chan bool)
	go func() {
		println("goroutine")
		ch <- true
	}()
	println("sec")
	<-ch
	close(ch)
	println("end")
}

func test17_2() {
	println("start")
	ch := make(chan int)
	var result int
	go func() {
		println("gouroutine1")
		var r int
		for i := 1; i <= 10; i++ {
			r += i
		}

		println(r)
		ch <-r
	}()
	go func() {
		println("gouroutine2")
		var r int
		for i := 1; i <= 10; i++ {
			r += i
		}
		println(r)
		ch <-r
	}()
	go func() {
		println("gouroutine3")

		ch <-11
	}()
	for i:=1; i<=3; i++{
		result += <-ch
	}
	println(result)
}

func test17_3()  {
	cond1 := make(chan int)
	cond2 := make(chan uint16)
	go func() {
		for i:=0; ; i++ {
			cond1 <- i
		}
	}()

	go func() {
		var i uint16
		for ; ; i++ {
			cond2 <- i
		}
	}()

	endCond := false
	for endCond != true{
		select {
		case a := <-cond1:
			if a>99 {
				println("end 1")
				endCond = true
			}
		case b := <-cond2:
			if b == 100{
				println("end 2")
				endCond = true
			}
		case <-time.After(time.Millisecond):
			println("end with timeout")
			endCond = true
		}
	}
	println("end")
}
