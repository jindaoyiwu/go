package main

import "fmt"

//阶乘
func main()  {
	func1()
}
func func1(){
	//输出结果
	//[ csdn baidu google] 6
	//前面多了几个空格，长度为6，与预期的结果不一致
	urls := make(map[string]string, 3)
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"
	fmt.Println(len(urls))
	names := make([]string, len(urls))
	for key, _ := range urls {
		names = append(names, key)
		fmt.Println(names,len(names))
	}
	fmt.Println(names,len(names))
}
func func2(){
	urls := make(map[string]string, 3)
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"
	names := make([]string, 0)
	for key, _ := range urls {
		names = append(names, key)
	}
	fmt.Println(names,len(names))
}
func func3(){
	urls := make(map[string]string, 3)
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"
	var names []string
	for key, _ := range urls {
		names = append(names, key)
	}
	fmt.Println(names,len(names))
}

func func4()  {
	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2)
	//fmt.Println(s1)
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}