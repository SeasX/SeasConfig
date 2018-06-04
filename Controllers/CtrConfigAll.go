package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type ConfigAllControl struct {
	GoCrab.Controller
}

func (ctr *ConfigAllControl) getParams() (pid string, aid string) {
	pid = ctr.Ctx.Input.Params[":PID"]
	aid = ctr.Ctx.Input.Params[":AID"]

	return pid, aid
}

//get all configs of an app
func (ctr *ConfigAllControl) Get() {
	pid, aid := ctr.getParams()

	if err := Models.ExistsAppConfigs(pid, aid); err != nil {
		ctr.RESTFaild(nil, err.Error())
	} else {
		ctr.RESTSuccess(Models.GetAppConfigs(pid, aid), nil)
	}
}
