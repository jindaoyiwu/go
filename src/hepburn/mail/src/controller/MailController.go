package controller

import (
	"hepburn/mail/service"
	"net/http"
	"strings"
)

func SendMail(w http.ResponseWriter, r *http.Request) service.Return {
	if err := r.ParseForm(); err != nil {
		return service.Return{
			Code:     400,
			Msg:      err.Error(),
			Data: map[string]interface{}{
				"error" : err,
			},
			Redirect: "",
		}
	}
	//定义收件人
	mailTo := strings.Split(r.PostFormValue("mailTo"), ",")
	//邮件主题为"Hello"
	subject := r.PostFormValue("subject")
	//邮件正文
	body := r.PostFormValue("body")
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string {
		"user": service.Env("mailFrom"),
		"pass": service.Env("mainPass"),
		"host": service.Env("mainHost"),
		"port": service.Env("mainPort"),
	}
	return service.SendMail(mailTo, subject, body, mailConn)
}