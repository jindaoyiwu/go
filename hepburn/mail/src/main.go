package main

import (
	"fmt"
	"hepburn/mail/controller"
	"hepburn/mail/service"
	"io"
	"log"
	"net/http"
)

// 自定义一个多路复用器的结构
type Mux struct{}

// 实现ServeHTTP方法
func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 判断URL并转到对应的handler处理
	if r.URL.Path == "/mail" {
		ret_json := service.JsonReturn(controller.SendMail(w, r))
		_, _ = io.WriteString(w, string(ret_json))
		return
	}
	http.NotFound(w, r)
	return
}


func main() {
	// 实例化一个自定义的路由器
	mux := Mux{}
	log.Println(service.Env("addr"))
	err := http.ListenAndServe(service.Env("addr"), mux)
	if err != nil {
		fmt.Println(err)
	}
}