package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
)

type ConfigControl struct {
	GoCrab.Controller
}

func (con *ConfigControl) getParams() (string, string, string) {
	pid := con.Ctx.Input.Params[":PID"]
	aid := con.Ctx.Input.Params[":AID"]
	key := con.Ctx.Input.Params[":KEY"]

	return pid, aid, key
}

//create new config
func (con *ConfigControl) Post() {
	pid, aid, key := con.getParams()
	GoCrab.Debug("Post PID", pid, "AID", aid, "KEY", key)
}

//get config
func (con *ConfigControl) Get() {
	pid, aid, key := con.getParams()
	GoCrab.Debug("Get PID", pid, "AID", aid, "KEY", key)
}

//update config value
func (con *ConfigControl) Put() {
	pid, aid, key := con.getParams()
	GoCrab.Debug("Put PID", pid, "AID", aid, "KEY", key)
}

//delete config by key
func (con *ConfigControl) Delete() {
	pid, aid, key := con.getParams()
	GoCrab.Debug("Delete PID", pid, "AID", aid, "KEY", key)
}

//exists config key
func (con *ConfigControl) Head() {
	pid, aid, key := con.getParams()
	GoCrab.Debug("Head PID", pid, "AID", aid, "KEY", key)
}
