package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
)

type ProductAllControl struct {
	GoCrab.Controller
}

//get all products info
func (con *ProductAllControl) Get() {
	GoCrab.Debug("Get ProductAllControl")
}
