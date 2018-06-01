package Models

import (
	"sync"
)

var (
	MapStore     map[string]map[string]map[string]interface{}
	MapStoreLock = new(sync.Mutex)
)

func init() {
	MapStore = make(map[string]map[string]map[string]interface{})
}

func GetAllStore() (maps map[string]map[string]map[string]interface{}) {
	return MapStore
}
