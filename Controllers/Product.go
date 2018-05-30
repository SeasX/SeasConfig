package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
)

type ProductControl struct {
	GoCrab.Controller
}

//get all input params in uri
func (con *ProductControl) getParams() string {
	pid := con.Ctx.Input.Params[":PID"]
	return pid
}

//create new product
func (con *ProductControl) Post() {
	pid := con.getParams()
	GoCrab.Debug("Post PID", pid)
}

//get product name
func (con *ProductControl) Get() {
	pid := con.getParams()
	GoCrab.Debug("Get PID", pid)
}

//update product info
func (con *ProductControl) Put() {
	pid := con.getParams()
	GoCrab.Debug("Put PID", pid)
}

//delete product by pid
func (con *ProductControl) Delete() {
	pid := con.getParams()
	GoCrab.Debug("Delete PID", pid)
}

//exists product by pid
func (con *ProductControl) Head() {
	pid := con.getParams()
	GoCrab.Debug("Head PID", pid)
}
