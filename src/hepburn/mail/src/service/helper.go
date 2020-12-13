package service

import "encoding/json"


type Return struct{
	Code int
	Msg string
	Data map[string]interface{}
	Redirect string
}
func JsonReturn(d Return) string {
	jsonReturn,_ :=  json.Marshal(d)
	return string(jsonReturn)
}
