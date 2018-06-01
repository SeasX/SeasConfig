package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Models"
)

type ProductAllControl struct {
	GoCrab.Controller
}

//get all products info
func (con *ProductAllControl) Get() {
	if !Models.ExistsProducts() {
		con.RESTFaild(nil, "Have No Product In SeasConfig")
	} else {
		con.RESTSuccess(Models.GetAllProducts(), nil)
	}
}
