# SeasConfig
配置管理中心


## 获得不带内存信息的状态页
 - http://127.0.0.1:6699/status
```
{
	AppName: "SeasConfig",
	Version: "0.1.0",
	Status: "OK",
	RuntimeGC: 30,
	StartTime: 1527084293,
	Duration: 334
}
```

## 获得带内存信息的状态页
 - http://127.0.0.1:6699/status?getMemStats=true
```
{
AppName: "SeasConfig",
Version: "0.1.0",
Status: "OK",
RuntimeGC: 30,
StartTime: 1527084150,
Duration: 18,
MemStats: {
	Alloc: 1020632,
	TotalAlloc: 1020632,
	Sys: 3084288,
	Lookups: 15,
	Mallocs: 6579,
	Frees: 334,
	HeapAlloc: 1020632,
	HeapSys: 1703936,
	HeapIdle: 98304,
	HeapInuse: 1605632,
	HeapReleased: 0,
	HeapObjects: 6245,
	StackInuse: 393216,
	StackSys: 393216,
	MSpanInuse: 19912,
	MSpanSys: 32768,
	MCacheInuse: 4800,
	MCacheSys: 16384,
	BuckHashSys: 2755,
	GCSys: 137216,
	OtherSys: 798013,
	NextGC: 4473924,
	LastGC: 0,
	PauseTotalNs: 0,
	NumGC: 0,
	NumForcedGC: 0,
	GCCPUFraction: 0,
	EnableGC: true,
	DebugGC: false
	}
}
```