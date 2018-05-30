package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
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
	GoCrab.Debug("Get AppAllControl PID", pid)
}
