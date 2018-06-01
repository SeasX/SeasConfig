package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type StatusControl struct {
	GoCrab.Controller
}

func (con *StatusControl) Get() {
	mem, _ := con.GetBool("getMemStats")

	if mem {
		obsWithM := Models.GetStatusWithMem()
		con.Data["json"] = obsWithM
	} else {
		obsWithOutM := Models.GetStatusWithOutMem()
		con.Data["json"] = obsWithOutM
	}

	con.ServeJson()
}
