package Controllers

import (
	"fmt"
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"time"
)

type AppControl struct {
	GoCrab.Controller
}

//get all params
func (con *AppControl) getParams() (string, string) {
	pid := con.Ctx.Input.Params[":PID"]
	aid := con.Ctx.Input.Params[":AID"]

	return pid, aid
}

//create new app
func (con *AppControl) Post() {
	pid, aid := con.getParams()
	GoCrab.Debug("Post PID", pid, "AID", aid)
}

//get app name
func (con *AppControl) Get() {
	pid, aid := con.getParams()
	GoCrab.Debug("Get PID", pid, "AID", aid)

	start := time.Now().UnixNano()
	fmt.Println("testLevelDB time start", start)

	a := make(map[string]map[string]interface{})
	c := make(map[string]interface{})
	a[pid] = c
	for i := 0; i < 1000000; i++ {
		c["a"] = "a"
	}

	fmt.Println(a)

	end := time.Now().UnixNano()
	fmt.Println("testLevelDB time end", end)
	fmt.Println("testLevelDB time total", (end-start)/1000/1000)

	con.Data["json"] = map[string]string{"Msg": "faild"}
	con.ServeJson()
}

//update app info
func (con *AppControl) Put() {
	pid, aid := con.getParams()
	GoCrab.Debug("Put PID", pid, "AID", aid)
}

//delete app by aid
func (con *AppControl) Delete() {
	pid, aid := con.getParams()
	GoCrab.Debug("Delete PID", pid, "AID", aid)
}

//exists app by aid
func (con *AppControl) Head() {
	pid, aid := con.getParams()
	GoCrab.Debug("Head PID", pid, "AID", aid)
}
