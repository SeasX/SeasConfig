package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type ProductControl struct {
	GoCrab.Controller
}

//get all input params in uri
func (con *ProductControl) getParams(getExcess bool) (pid string, name string, description string) {
	pid = con.Ctx.Input.Params[":PID"]
	if getExcess {
		name = con.GetString("name")
		description = con.GetString("description")
	} else {
		name = ""
		description = ""
	}

	return pid, name, description
}

//create new product
func (con *ProductControl) Post() {
	pid, name, description := con.getParams(true)
	Models.PutProduct(pid, name, description)

	con.RESTSuccess(pid, "Create Product Successful")
}

//get product name
func (con *ProductControl) Get() {
	pid, _, _ := con.getParams(false)
	product, have := Models.GetProduct(pid)
	if !have {
		con.RESTFaild(pid, "Can not find product")
	} else {
		con.RESTSuccess(product, nil)
	}
}

//update product info
func (con *ProductControl) Put() {
	pid, name, description := con.getParams(true)
	Models.PutProduct(pid, name, description)

	con.RESTSuccess(pid, "Update Product Successful")
}

//delete product by pid
func (con *ProductControl) Delete() {
	pid, _, _ := con.getParams(false)

	if !Models.ExistsProduct(pid) {
		con.RESTFaild(pid, "Can not find product")
	} else {
		if Models.DeleteProduct(pid) {
			con.RESTSuccess(pid, "Delete Product Successful")
		} else {
			con.RESTFaild(pid, "Delete Product unsuccessful")
		}
	}
}

//exists product by pid
func (con *ProductControl) Head() {
	pid, _, _ := con.getParams(false)

	if Models.ExistsProduct(pid) {
		con.RESTHeadSuccess()
	} else {
		con.RESTHeadNotFound()
	}
}
