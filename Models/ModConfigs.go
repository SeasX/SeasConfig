package Models

import (
	"github.com/SeasX/SeasConfig/Enums"
	"sync"
	//	"time"
	//	"github.com/CloudWise-OpenSource/GoCrab/Api"
)

type ConfigChan struct {
	Channel chan int
	Lock    *sync.Mutex
}

var (
	ConfigChans map[string]*ConfigChan
	ConfigStore map[string]map[string]map[string]interface{}
)

func init() {
	ConfigChans = make(map[string]*ConfigChan)
	ConfigStore = make(map[string]map[string]map[string]interface{})
}

func GetAllConfigs() (maps map[string]map[string]map[string]interface{}) {
	return ConfigStore
}

func GetProductConfigs(pid string) (maps map[string]map[string]interface{}) {
	return ConfigStore[pid]
}

func ExistsAppConfigs(pid string, aid string) error {
	if _, ok := ConfigStore[pid]; !ok {
		return Enums.NOT_FOUND_PRODUCT
	}

	if _, ok := ConfigStore[pid][aid]; !ok {
		return Enums.NOT_FOUND_APP
	}

	return nil
}

func GetAppConfigs(pid string, aid string) (maps map[string]interface{}) {
	return ConfigStore[pid][aid]
}

func InitConfigChan(pid string, aid string) bool {
	if _, ok := ConfigChans[pid+aid]; !ok {
		chanTmp := make(chan int)
		lockTmp := new(sync.Mutex)
		var configChan = ConfigChan{chanTmp, lockTmp}
		ConfigChans[pid+aid] = &configChan

		return true
	}

	return false
}

func ExistsConfig(pid string, aid string, key string) error {
	if _, ok := ConfigStore[pid]; !ok {
		return Enums.NOT_FOUND_PRODUCT
	}

	if _, ok := ConfigStore[pid][aid]; !ok {
		return Enums.NOT_FOUND_APP
	}

	if _, ok := ConfigStore[pid][aid][key]; !ok {
		return Enums.NOT_FOUND_CONFIG
	} else {
		return nil
	}
}

func GetConfig(pid string, aid string, key string) (value interface{}, err error) {
	if err := ExistsConfig(pid, aid, key); err != nil {
		return nil, err
	}

	if value, ok := ConfigStore[pid][aid][key]; !ok {
		return nil, Enums.NOT_FOUND_CONFIG
	} else {
		return value, nil
	}
}

func PutConfig(pid string, aid string, key string, value interface{}) bool {
	var have bool
	if _, ok := ConfigStore[pid]; !ok {
		ConfigStore[pid] = make(map[string]map[string]interface{})
	}

	if _, ok := ConfigStore[pid][aid]; !ok {
		tmpApp := make(map[string]interface{})
		ConfigStore[pid][aid] = tmpApp
		have = true
	}

	if !have {
		if ConfigStore[pid][aid][key] == value {
			return true
		}
	}

	ConfigChans[pid+aid].Lock.Lock()
	ConfigStore[pid][aid][key] = value
	ConfigChans[pid+aid].Lock.Unlock()

	ConfigChans[pid+aid].Channel <- 1

	return true
}

func DeleteConfig(pid string, aid string, key string) (bool, error) {
	if err := ExistsConfig(pid, aid, key); err != nil {
		return false, err
	}

	ConfigChans[pid+aid].Lock.Lock()
	delete(ConfigStore[pid][aid], key)
	ConfigChans[pid+aid].Lock.Unlock()

	ConfigChans[pid+aid].Channel <- 1
	return true, nil
}
