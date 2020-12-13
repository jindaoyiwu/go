package const_test

import "testing"

const (
	MONDAY = iota + 1
	TUESDAY
	WEDNESDAY
)
func TestConst(t *testing.T)  {
	t.Log(MONDAY, TUESDAY, WEDNESDAY)
}
