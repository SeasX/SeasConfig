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
func (con *AppControl) getParams(getExcess bool) (pid string, aid string, name string, description string) {
	pid = con.Ctx.Input.Params[":PID"]
	aid = con.Ctx.Input.Params[":AID"]

	if getExcess {
		name = con.GetString("name")
		description = con.GetString("description")
	} else {
		name = ""
		description = ""
	}

	return pid, aid, name, description
}

//create new app
func (con *AppControl) Post() {
	pid, aid, name, description := con.getParams(true)

	have := Models.ExistsProduct(pid)
	if !have {
		con.RESTFaild(pid, "Can not find product")
		return
	}

	if Models.InitAppChan(pid) {
		Tasks.WatchSaveApps(pid)
	}

	if Models.PutApp(pid, aid, name, description) {
		con.RESTSuccess(pid+"/"+aid, "Create App Successful")
	} else {
		con.RESTFaild(pid+"/"+aid, "Create App Faild")
	}
}

//get app name
func (con *AppControl) Get() {
	pid, aid, _, _ := con.getParams(false)

	have := Models.ExistsProduct(pid)
	if !have {
		con.RESTFaild(pid, "Can not find product")
		return
	}

	app, have := Models.GetApp(pid, aid)
	if !have {
		con.RESTFaild(pid+"/"+aid, "Can not find App")
	} else {
		con.RESTSuccess(app, nil)
	}
}

//update app info
func (con *AppControl) Put() {
	pid, aid, name, description := con.getParams(true)

	have := Models.ExistsProduct(pid)
	if !have {
		con.RESTFaild(pid, "Can not find product")
		return
	}

	if Models.InitAppChan(pid) {
		Tasks.WatchSaveApps(pid)
	}

	if Models.PutApp(pid, aid, name, description) {
		con.RESTSuccess(pid+"/"+aid, "Update App Successful")
	} else {
		con.RESTFaild(pid+"/"+aid, "Update App Faild")
	}
}

//delete app by aid
func (con *AppControl) Delete() {
	pid, aid, _, _ := con.getParams(false)

	have := Models.ExistsProduct(pid)
	if !have {
		con.RESTFaild(pid, "Can not find product")
		return
	}

	have = Models.ExistsApp(pid, aid)
	if !have {
		con.RESTFaild(pid+"/"+aid, "Can not find App")
		return
	}

	if Models.InitAppChan(pid) {
		Tasks.WatchSaveApps(pid)
	}

	if Models.DeleteApp(pid, aid) {
		con.RESTSuccess(pid+"/"+aid, "Delete App Successful")
	} else {
		con.RESTFaild(pid+"/"+aid, "Delete App Faild")
	}
}

//exists app by aid
func (con *AppControl) Head() {
	pid, aid, _, _ := con.getParams(false)

	if Models.ExistsApp(pid, aid) {
		con.RESTHeadSuccess()
	} else {
		con.RESTHeadNotFound()
	}
}
