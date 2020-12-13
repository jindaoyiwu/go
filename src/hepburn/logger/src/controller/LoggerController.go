package controller

import (
	"encoding/json"
	"hepburn/logger/service"
	"log"
	"net/http"
)

func LoggerLevel(w http.ResponseWriter,  r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		log.Println("表单解析错误,",err)
		return
	}
	method := r.PostFormValue("method")
	dataLine := r.PostFormValue("data[line]")
	dataFile := r.PostFormValue("data[file]")
	dataMsg := r.PostFormValue("data[msg]")
	dataValues := r.PostForm["data[values]"]
	temp := make(map[string]interface{})
	temp["msg"] = dataMsg
	temp["file"] = dataFile
	temp["line"] = dataLine
	temp["values"] = dataValues
	tempV,_ := json.Marshal(temp)
	switch method {
		case "Debug":
			service.Debug(string(tempV))
		case "Info":
			service.Info(string(tempV))
		case "Warn":
			service.Warn(string(tempV))
		case "Error":
			//如果是错误就发送邮件
			data := make(map[string]interface{})
			data["subject"] = "报警邮件"
			data["body"] = string(tempV)
			data["mailTo"] = service.Env("mailTo")
			//异步发送邮件
			go service.SendEmail(service.Env("mailUrl"), data)
			service.Error(string(tempV))
		case "Panic":
			service.Panic(string(tempV))
		case "Fatal":
			service.Fatal(string(tempV))
	}
}