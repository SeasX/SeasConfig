package main

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Controllers"
	_ "github.com/SeasX/SeasConfig/Tasks"
)

func main() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			GoCrab.Error(panicErr)
		}
	}()

	GoCrab.SetLogger("file", `{"filename":"logs/error.log"}`)

	//状态页
	GoCrab.RESTRouter("/status", &Controllers.StatusControl{})

	//产品
	GoCrab.Router("/product/all", &Controllers.ProductAllControl{})
	GoCrab.Router("/product/:PID", &Controllers.ProductControl{})

	//产品下应用
	GoCrab.Router("/app/:PID/all", &Controllers.AppAllControl{})
	GoCrab.Router("/app/:PID/:AID", &Controllers.AppControl{})

	//应用的具体配置
	GoCrab.Router("/config/:PID/:AID/all", &Controllers.ConfigAllControl{})
	GoCrab.Router("/config/:PID/:AID/:KEY", &Controllers.ConfigControl{})

	GoCrab.Run()
}
