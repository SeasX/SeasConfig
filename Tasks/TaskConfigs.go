package Tasks

import (
	"encoding/gob"
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/CloudWise-OpenSource/GoCrab/Helpers"
	"github.com/SeasX/SeasConfig/Enums"
	"github.com/SeasX/SeasConfig/Models"
	"os"
)

func LoadConfigs() error {
	GoCrab.Debug("LoadConfigs Start")

	for pid, apps := range Models.Apps {
		if _, ok := Models.ConfigStore[pid]; !ok {
			Models.ConfigStore[pid] = make(map[string]map[string]interface{})
		}

		if len(apps) > 0 {
			for aid, app := range apps {
				if _, ok := Models.ConfigStore[pid][aid]; !ok {
					configs := make(map[string]interface{})
					GoCrab.Debug("LoadConfig", app.Name, "start.")
					GoCrab.Debug("App.Name ->", app.Name, "App.CreateTime ->", app.CreateTime)
					LoadConfig(pid, aid, &configs)
					Models.ConfigStore[pid][aid] = configs
					GoCrab.Debug("LoadConfig", app.Name, "end.")
				}
			}
		}
	}

	GoCrab.Debug("LoadConfigs Complete")

	return nil
}

func LoadConfig(pid string, aid string, apps interface{}) error {
	fileName := Enums.GetDBConfigFile(pid, aid)

	if !Helpers.FileExists(fileName) {
		GoCrab.Warn("LoadConfig Do not exists", fileName)
		return nil
	}

	file, err := os.Open(fileName)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(apps)
	}

	if err != nil {
		GoCrab.Error("LoadConfig Error", fileName, err)
	} else {
		GoCrab.Debug("LoadConfig Successful", fileName)
	}

	file.Close()

	return err
}

func SaveConfigs(pid string, aid string, apps interface{}) error {
	dir := Enums.GetDBProductDir(pid)

	if !Helpers.FileExists(dir) {
		if _, err := Helpers.MakeDir(dir); err != nil {
			GoCrab.Error("MakeDir error", dir, err)
			return err
		}
	}

	file, err := os.Create(Enums.GetDBConfigFile(pid, aid))
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(apps)
	} else {
		GoCrab.Error("SaveConfigs error", Enums.GetDBConfigFile(pid, aid), err)
	}
	file.Close()

	return err
}
