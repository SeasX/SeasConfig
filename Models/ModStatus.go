package Models

import (
	"github.com/SeasX/SeasConfig/Enums"
	"runtime"
	"time"
)

type mStats struct {
	Alloc         uint64
	TotalAlloc    uint64
	Sys           uint64
	Lookups       uint64
	Mallocs       uint64
	Frees         uint64
	HeapAlloc     uint64
	HeapSys       uint64
	HeapIdle      uint64
	HeapInuse     uint64
	HeapReleased  uint64
	HeapObjects   uint64
	StackInuse    uint64
	StackSys      uint64
	MSpanInuse    uint64
	MSpanSys      uint64
	MCacheInuse   uint64
	MCacheSys     uint64
	BuckHashSys   uint64
	GCSys         uint64
	OtherSys      uint64
	NextGC        uint64
	LastGC        uint64
	PauseTotalNs  uint64
	NumGC         uint32
	NumForcedGC   uint32
	GCCPUFraction float64
	EnableGC      bool
	DebugGC       bool
}

var (
	StatusTM statusM
	StatusT  status
	sysM     runtime.MemStats
	thisM    mStats
)

type statusM struct {
	AppName   string
	Version   string
	Status    string
	RuntimeGC int
	StartTime int64
	Duration  int64
	MemStats  mStats
}
type status struct {
	AppName   string
	Version   string
	Status    string
	RuntimeGC int
	StartTime int64
	Duration  int64
}

func init() {
	appName := Enums.APP_NAME
	version := Enums.APP_VERSION
	state := Enums.APP_STATUS_OK

	startTime := time.Now().Unix()
	runtimeGC := Enums.GetRuntimeGC()

	StatusTM = statusM{appName, version, state, runtimeGC, startTime, 0, thisM}
	StatusT = status{appName, version, state, runtimeGC, startTime, 0}
}

func GetStatusWithMem() statusM {
	now := time.Now().Unix() - StatusT.StartTime
	StatusTM.Duration = now
	runtime.ReadMemStats(&sysM)

	thisM.Alloc = sysM.Alloc
	thisM.TotalAlloc = sysM.TotalAlloc
	thisM.Sys = sysM.Sys
	thisM.Lookups = sysM.Lookups
	thisM.Mallocs = sysM.Mallocs
	thisM.Frees = sysM.Frees
	thisM.HeapAlloc = sysM.HeapAlloc
	thisM.HeapSys = sysM.HeapSys
	thisM.HeapIdle = sysM.HeapIdle
	thisM.HeapInuse = sysM.HeapInuse
	thisM.HeapReleased = sysM.HeapReleased
	thisM.HeapObjects = sysM.HeapObjects
	thisM.StackInuse = sysM.StackInuse
	thisM.StackSys = sysM.StackSys
	thisM.MSpanInuse = sysM.MSpanInuse
	thisM.MSpanSys = sysM.MSpanSys
	thisM.MCacheInuse = sysM.MCacheInuse
	thisM.MCacheSys = sysM.MCacheSys
	thisM.BuckHashSys = sysM.BuckHashSys
	thisM.GCSys = sysM.GCSys
	thisM.OtherSys = sysM.OtherSys
	thisM.NextGC = sysM.NextGC
	thisM.LastGC = sysM.LastGC
	thisM.PauseTotalNs = sysM.PauseTotalNs
	thisM.BuckHashSys = sysM.BuckHashSys
	thisM.NumGC = sysM.NumGC
	thisM.NumForcedGC = sysM.NumForcedGC
	thisM.GCCPUFraction = sysM.GCCPUFraction
	thisM.EnableGC = sysM.EnableGC
	thisM.DebugGC = sysM.DebugGC

	StatusTM.MemStats = thisM
	return StatusTM
}

func GetStatusWithOutMem() status {
	StatusT.Duration = time.Now().Unix() - StatusT.StartTime
	return StatusT
}
