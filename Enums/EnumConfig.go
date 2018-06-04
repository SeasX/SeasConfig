package Enums

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
)

var (
	Compress            string
	ChannelCount        int
	RuntimeGC           int
	ConnectTimeout      int
	ResponseTimeout     int
	MaxIdleConnsPerHost int
	IdleConnTimeout     int
	DBProductsFile      string
)

func GetDBProductsFile() string {
	if DBProductsFile == "" {
		DBProductsFile = DB_DIR + DB_FILE_PRODUCTS
	}
	return DBProductsFile
}

func GetDBProductDir(pid string) string {
	return DB_DIR + pid
}

func GetDBAppsFile(pid string) string {
	return DB_DIR + pid + "/" + DB_FILE_APPS
}

func GetDBConfigFile(pid string, aid string) string {
	return DB_DIR + pid + "/" + aid + "_" + DB_FILE_CONFIG
}

func CheckIfDebug() bool {
	return GoCrab.RunMode == GoCrab.RUNMODE_DEV
}

func GetChannelCount() int {
	if ChannelCount != 0 {
		return ChannelCount
	}

	if channelCount, err := GoCrab.AppConfig.Int("Channel"); err == nil {
		ChannelCount = channelCount
	} else {
		ChannelCount = ENUM_CHANNEL_COUNT
	}

	return ChannelCount
}

func GetRuntimeGC() int {
	if RuntimeGC != 0 {
		return RuntimeGC
	}

	if runtimeGC, err := GoCrab.AppConfig.Int("RuntimeGC"); err == nil {
		RuntimeGC = runtimeGC
	} else {
		RuntimeGC = RUNTIME_GC_TIME_DEFAULT
	}

	return RuntimeGC
}
