package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type ProductControl struct {
	GoCrab.Controller
}

//get all input params in uri
func (ctr *ProductControl) getParams(getExcess bool) (pid string, name string, description string) {
	pid = ctr.Ctx.Input.Params[":PID"]
	if getExcess {
		name = ctr.GetString("name")
		description = ctr.GetString("description")
	} else {
		name = ""
		description = ""
	}

	return pid, name, description
}

//create new product
func (ctr *ProductControl) Post() {
	pid, name, description := ctr.getParams(true)
	Models.PutProduct(pid, name, description)

	ctr.RESTSuccess(pid, "Create Product Successful")
}

//get product name
func (ctr *ProductControl) Get() {
	pid, _, _ := ctr.getParams(false)
	if product, err := Models.GetProduct(pid); err != nil {
		ctr.RESTFaild(pid, err.Error())
	} else {
		ctr.RESTSuccess(product, nil)
	}
}

//update product info
func (ctr *ProductControl) Put() {
	pid, name, description := ctr.getParams(true)
	Models.PutProduct(pid, name, description)

	ctr.RESTSuccess(pid, "Update Product Successful")
}

//delete product by pid
func (ctr *ProductControl) Delete() {
	pid, _, _ := ctr.getParams(false)

	if err := Models.ExistsProduct(pid); err != nil {
		ctr.RESTFaild(pid, err.Error())
	} else {
		if Models.DeleteProduct(pid) {
			ctr.RESTSuccess(pid, "Delete Product Successful")
		} else {
			ctr.RESTFaild(pid, "Delete Product unsuccessful")
		}
	}
}

//exists product by pid
func (ctr *ProductControl) Head() {
	pid, _, _ := ctr.getParams(false)

	if err := Models.ExistsProduct(pid); err != nil {
		ctr.RESTHeadNotFound()
	} else {
		ctr.RESTHeadSuccess()
	}
}
