package Tasks

import (
	"encoding/gob"
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/CloudWise-OpenSource/GoCrab/Helpers"
	"github.com/SeasX/SeasConfig/Enums"
	"github.com/SeasX/SeasConfig/Models"
	"os"
)

func LoadApps() error {
	GoCrab.Debug("LoadApps Start")

	for pid, _ := range Models.Products {
		app := make(map[string]*Models.App)
		LoadApp(pid, &app)
		Models.Apps[pid] = app
	}

	GoCrab.Debug("LoadApps Successful")

	return nil
}

func LoadApp(pid string, apps interface{}) error {
	fileName := Enums.GetDBAppsFile(pid)

	if !Helpers.FileExists(fileName) {
		GoCrab.Warn("LoadApp Do not exists", fileName)
		return nil
	}

	file, err := os.Open(fileName)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(apps)

		GoCrab.Debug("LoadApp Successful", fileName)
	} else {
		GoCrab.Error("LoadApp Error", fileName, err)
	}
	file.Close()

	return err
}

func SaveApps(pid string, apps interface{}) error {
	dir := Enums.GetDBProductDir(pid)

	if !Helpers.FileExists(dir) {
		if _, err := Helpers.MakeDir(dir); err != nil {
			GoCrab.Error("MakeDir error", dir, err)
			return err
		}
	}

	file, err := os.Create(Enums.GetDBAppsFile(pid))
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(apps)
	} else {
		GoCrab.Error("SaveApps error", Enums.GetDBAppsFile(pid), err)
	}
	file.Close()

	return err
}
