package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type baseController struct {
	beego.Controller
}

type DataResponse struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	ServerTime string      `json:"servertime"`
}

func Reponse(code int, data interface{}, msg string) DataResponse {
	resp := DataResponse{
		Code:       code,
		Msg:        msg,
		Data:       data,
		ServerTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	return resp
}

type Request struct {
	Code int
	Msg  string
	Data string
}
