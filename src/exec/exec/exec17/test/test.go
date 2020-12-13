package test

import "fmt"

var Config map[string]interface{}
func InitT()  {
	Config = make( map[string]interface{}, 1)
	Config["aa"] = &struct {
		ww string
	}{ww:"1234"}
	fmt.Println(Config)
}