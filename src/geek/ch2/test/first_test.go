package test

import (
	"fmt"
	"testing"
)
/**
文件要以_test结尾
方法要以Test开头
 */
func TestFirstTry(t *testing.T)  {
	fmt.Printf("test")
	t.Log("my first test")
}
