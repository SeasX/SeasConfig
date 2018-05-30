package Controllers

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
)

type ConfigAllControl struct {
	GoCrab.Controller
}

//get all configs of an app
func (con *ConfigAllControl) Get() {
	GoCrab.Debug("Get ConfigAllControl")
}
