# SeasConfig
配置管理中心


## 状态页
### 获得不带内存信息的状态页
 - curl http://127.0.0.1:6699/status
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

### 获得带内存信息的状态页
 - curl http://127.0.0.1:6699/status?getMemStats=true
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

## 产品API
### 获取所有产品概述
 - GET http://127.0.0.1:6699/product/all
```
curl http://127.0.0.1:6699/product/all

{
    "Code": 1000,
    "Data": {
        "test_product": {
            "Name": "product_name",
            "Description": "product_description",
            "CreateTime": 1527828786,
            "UpdateTime": 1527828786
        },
        "test_product_2": {
            "Name": "product_name",
            "Description": "product_description",
            "CreateTime": 1527833319,
            "UpdateTime": 1527833506
        }
    },
    "Message": "ok"
}
```

### 创建新产品
 - POST http://127.0.0.1:6699/product/{:PID}
```

curl -d "name=product_name&description=product_description" http://127.0.0.1:6699/product/test_product

{
    "Code": 1000,
    "Data": "test_product",
    "Message": "Create Product Successful"
}
```

### 更新原有产品
 - PUT http://127.0.0.1:6699/product/{:PID}
```
curl -X PUT -d "name=updated_name&description=updated_description" http://127.0.0.1:6699/product/test_product

{
    "Code": 1000,
    "Data": "test_product",
    "Message": "Update Product Successful"
}

```

### 获取某产品概述
 - GET http://127.0.0.1:6699/product/{:PID}
```
curl http://127.0.0.1:6699/product/test_product

{
    "Code": 1000,
    "Data": {
        "Name": "updated_name",
        "Description": "updated_description",
        "CreateTime": 1527828786,
        "UpdateTime": 1527834696
    },
    "Message": "ok"
}

```

### 删除某产品
 - DELETE http://127.0.0.1:6699/product/{:PID}
```
curl -X DELETE http://127.0.0.1:6699/product/test_product

{
    "Code": 1000,
    "Data": "test_product",
    "Message": "Delete Product Successful"
}
```

### 检查某产品是否存在
 - HEAD http://127.0.0.1:6699/product/{:PID}
```
curl -i -X HEAD http://127.0.0.1:6699/product/test_product

HTTP/1.1 200 OK
Date: Fri, 01 Jun 2018 07:03:39 GMT

```

## 应用API
### 获取某产品下所有应用概述
 - GET http://127.0.0.1:6699/app/{:PID}/all
```
curl http://127.0.0.1:6699/app/test_product/all

{
    "Code": 1000,
    "Data": {
        "test_app": {
            "Name": "updated_name",
            "Description": "updated_description",
            "CreateTime": 1527835739,
            "UpdateTime": 1527835739
        },
        "test_app_2": {
            "Name": "updated_name",
            "Description": "updated_description",
            "CreateTime": 1527835894,
            "UpdateTime": 1527835894
        }
    },
    "Message": "ok"
}

```

### 在某产品下创建一个应用
 - POST http://127.0.0.1:6699/app/{:PID}/{:AID}
```
curl -d "name=product_name&description=product_description" http://127.0.0.1:6699/app/test_product/test_app

{
    "Code": 1000,
    "Data": "test_product/test_app",
    "Message": "Create App Successful"
}
```

### 更新某应用
 - PUT http://127.0.0.1:6699/app/{:PID}/{:AID}
```
curl -X PUT -d "name=updated_name&description=updated_description" http://127.0.0.1:6699/app/test_product/test_app

{
    "Code": 1000,
    "Data": "test_product/test_app",
    "Message": "Update App Successful"
}
```

### 获取某应用概述
 - GET http://127.0.0.1:6699/app/{:PID}/{:AID}
```
curl http://127.0.0.1:6699/app/test_product/test_app

{
    "Code": 1000,
    "Data": {
        "Name": "updated_name",
        "Description": "updated_description",
        "CreateTime": 1527835739,
        "UpdateTime": 1527835739
    },
    "Message": "ok"
}
```

### 删除某应用
 - DELETE http://127.0.0.1:6699/app/{:PID}/{:AID}
```
curl -X DELETE http://127.0.0.1:6699/app/test_product/test_app

{
    "Code": 1000,
    "Data": "test_product/test_app",
    "Message": "Delete App Successful"
}
```

### 检查某应用是否存在
 - HEAD http://127.0.0.1:6699/app/{:PID}/{:AID}
```
curl -i -X HEAD http://127.0.0.1:6699/app/test_product/test_app

HTTP/1.1 404 Not Found
Date: Fri, 01 Jun 2018 07:06:39 GMT

```