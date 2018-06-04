package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
	"github.com/SeasX/SeasConfig/Tasks"
)

type ConfigControl struct {
	GoCrab.Controller
}

func (ctr *ConfigControl) getParams(getExcess bool) (pid string, aid string, key string, value string) {
	pid = ctr.Ctx.Input.Params[":PID"]
	aid = ctr.Ctx.Input.Params[":AID"]
	key = ctr.Ctx.Input.Params[":KEY"]

	if getExcess {
		value = ctr.GetString("value")
	} else {
		value = ""
	}

	return pid, aid, key, value
}

func (ctr *ConfigControl) validChan(pid string, aid string) {
	if Models.InitConfigChan(pid, aid) {
		Tasks.WatchSaveConfigs(pid, aid)
	}
}

//create new config
func (ctr *ConfigControl) Post() {
	pid, aid, key, value := ctr.getParams(true)

	ctr.validChan(pid, aid)

	if Models.PutConfig(pid, aid, key, value) {
		ctr.RESTSuccess(pid+"/"+aid+"/"+key, "Create Config Successful")
	} else {
		ctr.RESTFaild(pid+"/"+aid+"/"+key, "Create Config Faild")
	}
}

//get config
func (ctr *ConfigControl) Get() {
	pid, aid, key, _ := ctr.getParams(false)

	if value, err := Models.GetConfig(pid, aid, key); err != nil {
		ctr.RESTFaild(pid+"/"+aid+"/"+key, err.Error())
	} else {
		ctr.RESTSuccess(value, nil)
	}
}

//update config value
func (ctr *ConfigControl) Put() {
	pid, aid, key, value := ctr.getParams(true)

	ctr.validChan(pid, aid)

	if Models.PutConfig(pid, aid, key, value) {
		ctr.RESTSuccess(pid+"/"+aid+"/"+key, "Update Config Successful")
	} else {
		ctr.RESTFaild(pid+"/"+aid+"/"+key, "Update Config Faild")
	}
}

//delete config by key
func (ctr *ConfigControl) Delete() {
	pid, aid, key, _ := ctr.getParams(false)

	ctr.validChan(pid, aid)

	if _, err := Models.DeleteConfig(pid, aid, key); err != nil {
		ctr.RESTFaild(pid+"/"+aid+"/"+key, err.Error())
	} else {
		ctr.RESTSuccess(pid+"/"+aid+"/"+key, "Delete Config Successful")
	}
}

//exists config key
func (ctr *ConfigControl) Head() {
	pid, aid, key, _ := ctr.getParams(false)

	if err := Models.ExistsConfig(pid, aid, key); err != nil {
		ctr.RESTHeadNotFound()
	} else {
		ctr.RESTHeadSuccess()
	}
}
