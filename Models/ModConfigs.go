package Models

func GetProductConfigs(pid string) (maps map[string]map[string]interface{}) {
	return MapStore[pid]
}

func GetAppConfigs(pid string, aid string) (maps map[string]interface{}) {

	return MapStore[pid][aid]
}

func PutConfig(pid string, aid string, key string, value interface{}) bool {
	return true
}

func GetConfig(pid string, aid string, key string) (value interface{}, err error) {

	return value, err
}

func DeleteConfig(pid string, aid string, key string) bool {
	return true
}
