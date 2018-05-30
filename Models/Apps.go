package Models

import (
	"sync"
	"time"
)

type App struct {
	Name        string
	Description string
	CreateTime  int64
	UpdateTime  int64
}

var (
	Apps    map[string]*App
	AppLock map[string]*sync.Mutex
)

func init() {
	Apps = make(map[string]*App)
	AppLock = make(map[string]*sync.Mutex)
}

func GetAllApps() (maps map[string]*App) {
	return Apps
}

func GetApp(aid string) (app *App) {
	return Apps[aid]
}

func PutApp(pid string, aid string, name string, descr string) bool {
	if _, ok := AppLock[aid]; !ok {
		AppLock[aid] = new(sync.Mutex)
	}
	AppLock[aid].Lock()
	defer AppLock[aid].Unlock()

	if _, ok := Apps[aid]; !ok {
		Apps[aid] = new(App)
		Apps[aid].CreateTime = time.Now().Unix()
	}

	Apps[aid].Name = name
	Apps[aid].Description = descr
	Apps[aid].UpdateTime = time.Now().Unix()

	return true
}

func DeleteApp(aid string) bool {
	AppLock[aid].Lock()
	defer AppLock[aid].Unlock()

	delete(Apps, aid)

	return true
}
