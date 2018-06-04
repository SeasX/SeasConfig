package Enums

import (
	"errors"
)

var (
	HAVE_NO_PRODUCT = "Have No Product In SeasConfig"

	NOT_FOUND_PRODUCT = errors.New("Can Not Found Product")
	NOT_FOUND_APP     = errors.New("Can Not Found App")
	NOT_FOUND_CONFIG  = errors.New("Can Not Found Config")
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
