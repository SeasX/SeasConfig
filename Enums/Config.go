package Enums

import (
	"github.com/CloudWise-OpenSource/GoCrab/Api"
)

var (
	ConfigEnv       string
	DefaultApi      string
	BatchDefaultApi string
	Compress        string
	ChannelCount    int
	UdpPort         int
	QueueLenMax     int
	GziperCount     int

	RuntimeGC int

	ConnectTimeout      int
	ResponseTimeout     int
	MaxIdleConnsPerHost int
	IdleConnTimeout     int

	AssignSendUrl      string
	AssignBatchSendUrl string

	AssignRegisterHost string

	QueueLenExceedStrategy string
)

func CheckIfDebug() bool {
	return GoCrab.RunMode == GoCrab.RUNMODE_DEV
}

func GetEnv() string {
	if ConfigEnv != "" {
		return ConfigEnv
	}

	if setEnv := GoCrab.AppConfig.String("SetEnv"); setEnv != "" {
		ConfigEnv = setEnv
	}

	return ConfigEnv
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
