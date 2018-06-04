package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/SeasX/SeasConfig/Enums"
	"github.com/SeasX/SeasConfig/Models"
)

type ProductAllControl struct {
	GoCrab.Controller
}

//get all products info
func (ctr *ProductAllControl) Get() {
	if !Models.ExistsProducts() {
		ctr.RESTFaild(nil, Enums.HAVE_NO_PRODUCT)
	} else {
		ctr.RESTSuccess(Models.GetAllProducts(), nil)
	}
}
