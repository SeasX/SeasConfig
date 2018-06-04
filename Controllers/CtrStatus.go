package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type StatusControl struct {
	GoCrab.Controller
}

func (ctr *StatusControl) Get() {
	mem, _ := ctr.GetBool("getMemStats")

	if mem {
		obsWithM := Models.GetStatusWithMem()
		ctr.Data["json"] = obsWithM
	} else {
		obsWithOutM := Models.GetStatusWithOutMem()
		ctr.Data["json"] = obsWithOutM
	}

	ctr.ServeJson()
}
