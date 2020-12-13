package ch3

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	t.Log(len(s))
	if s == ""{
		fmt.Println("s 是空字符")
	}
}