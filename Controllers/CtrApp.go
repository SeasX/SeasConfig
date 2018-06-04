package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
	"github.com/SeasX/SeasConfig/Tasks"
)

type AppControl struct {
	GoCrab.Controller
}

//get all params
func (ctr *AppControl) getParams(getExcess bool) (pid string, aid string, name string, description string) {
	pid = ctr.Ctx.Input.Params[":PID"]
	aid = ctr.Ctx.Input.Params[":AID"]

	if getExcess {
		name = ctr.GetString("name")
		description = ctr.GetString("description")
	} else {
		name = ""
		description = ""
	}

	return pid, aid, name, description
}

func (ctr *AppControl) validChan(pid string) {
	if Models.InitAppChan(pid) {
		Tasks.WatchSaveApps(pid)
	}
}

//create new app
func (ctr *AppControl) Post() {
	pid, aid, name, description := ctr.getParams(true)

	if err := Models.ExistsProduct(pid); err != nil {
		ctr.RESTFaild(pid, err.Error())
		return
	}

	ctr.validChan(pid)

	if Models.PutApp(pid, aid, name, description) {
		ctr.RESTSuccess(pid+"/"+aid, "Create App Successful")
	} else {
		ctr.RESTFaild(pid+"/"+aid, "Create App Faild")
	}
}

//get app name
func (ctr *AppControl) Get() {
	pid, aid, _, _ := ctr.getParams(false)

	if err := Models.ExistsProduct(pid); err != nil {
		ctr.RESTFaild(pid, err.Error())
		return
	}

	if app, err := Models.GetApp(pid, aid); err != nil {
		ctr.RESTFaild(pid+"/"+aid, err.Error())
	} else {
		ctr.RESTSuccess(app, nil)
	}
}

//update app info
func (ctr *AppControl) Put() {
	pid, aid, name, description := ctr.getParams(true)

	if err := Models.ExistsProduct(pid); err != nil {
		ctr.RESTFaild(pid, err.Error())
		return
	}

	ctr.validChan(pid)

	if Models.PutApp(pid, aid, name, description) {
		ctr.RESTSuccess(pid+"/"+aid, "Update App Successful")
	} else {
		ctr.RESTFaild(pid+"/"+aid, "Update App Faild")
	}
}

//delete app by aid
func (ctr *AppControl) Delete() {
	pid, aid, _, _ := ctr.getParams(false)

	if err := Models.ExistsProduct(pid); err != nil {
		ctr.RESTFaild(pid, err.Error())
		return
	}

	if err := Models.ExistsApp(pid, aid); err != nil {
		ctr.RESTFaild(pid+"/"+aid, err.Error())
		return
	}

	ctr.validChan(pid)

	if Models.DeleteApp(pid, aid) {
		ctr.RESTSuccess(pid+"/"+aid, "Delete App Successful")
	} else {
		ctr.RESTFaild(pid+"/"+aid, "Delete App Faild")
	}
}

//exists app by aid
func (ctr *AppControl) Head() {
	pid, aid, _, _ := ctr.getParams(false)

	if err := Models.ExistsApp(pid, aid); err != nil {
		ctr.RESTHeadNotFound()
	} else {
		ctr.RESTHeadSuccess()
	}
}
