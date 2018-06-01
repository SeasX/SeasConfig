package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type AppAllControl struct {
	GoCrab.Controller
}

//get all params
func (con *AppAllControl) getParams() string {
	pid := con.Ctx.Input.Params[":PID"]

	return pid
}

//get all app list of an product
func (con *AppAllControl) Get() {
	pid := con.getParams()

	if !Models.ExistsApps(pid) {
		con.RESTFaild(nil, "Have No Apps In This Product")
	} else {
		con.RESTSuccess(Models.GetAllApps(pid), nil)
	}
}
