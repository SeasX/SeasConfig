package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type AppAllControl struct {
	GoCrab.Controller
}

//get all params
func (ctr *AppAllControl) getParams() string {
	pid := ctr.Ctx.Input.Params[":PID"]

	return pid
}

//get all app list of an product
func (ctr *AppAllControl) Get() {
	pid := ctr.getParams()

	if !Models.ExistsApps(pid) {
		ctr.RESTFaild(nil, "Have No Apps In This Product")
	} else {
		ctr.RESTSuccess(Models.GetAllApps(pid), nil)
	}
}
