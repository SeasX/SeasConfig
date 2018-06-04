package Tasks

import (
	"encoding/gob"
	"github.com/CloudWise-OpenSource/GoCrab/Api"
	"github.com/CloudWise-OpenSource/GoCrab/Helpers"
	"github.com/SeasX/SeasConfig/Enums"
	"os"
)

func LoadProducts(products interface{}) error {
	GoCrab.Debug("LoadProducts Start")

	productDbFile := Enums.GetDBProductsFile()

	if !Helpers.FileExists(productDbFile) {
		GoCrab.Warn("LoadProducts Do not exists", productDbFile)
		return nil
	}

	file, err := os.Open(productDbFile)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(products)
	}

	if err != nil {
		GoCrab.Error("LoadProducts Error", productDbFile, err)
	} else {
		GoCrab.Debug("LoadProducts Complete")
	}

	file.Close()

	return err
}

func SaveProducts(products interface{}) error {
	file, err := os.Create(Enums.GetDBProductsFile())
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(products)
		GoCrab.Debug("SaveProducts Successful")
	} else {
		GoCrab.Error("SaveProducts Error", err)
	}
	file.Close()

	return err
}
