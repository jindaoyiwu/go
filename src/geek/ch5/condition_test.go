package ch5

import (
	"runtime"
	"testing"
)

func TestLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestIf(t *testing.T) {
	if v, err := someFun(); err != "nil" {
		t.Log(v)
	} else {
		t.Log("is a err")
	}
}

func someFun() (int, string) {
	return 1, ""
}

func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("os.x")
	case "linux":
		t.Log("linux")
	default:
		t.Log("other")

	}
}

func TestMutiCase(t *testing.T)  {
	for i:=0;i<5 ;i++  {
		switch i {
		case 1,2:
			t.Log("1,2")
		case 3,4:
			t.Log(3,4)
		default:
			t.Log("other")

		}
	}
}