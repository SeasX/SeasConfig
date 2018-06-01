package Enums

var (
	JSON_FAILD = map[string]interface{}{"Code": 999, "Message": "faild", "Data": ""}
)

const APP_NAME = "SeasConfig"
const APP_VERSION = "0.1.0"
const APP_STATUS_OK = "OK"

const API_VERSION = "v2"
const API_VERSION_V3 = "v3"

const ENUM_CHANNEL_COUNT = 500

const RUNTIME_GC_TIME_DEFAULT = 1

const DB_DIR = "./db/"
const DB_FILE_PRODUCTS = "products.gob"
const DB_FILE_APPS = "apps.gob"
const DB_FILE_CONFIG = "config.gob"
