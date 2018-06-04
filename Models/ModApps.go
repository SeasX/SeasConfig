package Models

import (
	"github.com/CloudWise-OpenSource/GoCrab/Helpers"
	"github.com/SeasX/SeasConfig/Enums"
	"sync"
	"time"
)

type App struct {
	Name        string
	Description string
	CreateTime  int64
	UpdateTime  int64
}

type AppChan struct {
	Channel chan int
	Lock    *sync.Mutex
}

var (
	Apps     map[string]map[string]*App
	AppChans map[string]*AppChan
)

func init() {
	Apps = make(map[string]map[string]*App)
	AppChans = make(map[string]*AppChan)
}

func GetAllApps(pid string) (maps map[string]*App) {
	return Apps[pid]
}

func ExistsApps(pid string) (have bool) {
	if len(Apps[pid]) < 1 {
		return false
	}

	return true
}

func GetApp(pid string, aid string) (app *App, err error) {
	if _, ok := Apps[pid][aid]; !ok {
		return nil, Enums.NOT_FOUND_APP
	}

	return Apps[pid][aid], nil
}

func ExistsApp(pid string, aid string) error {
	if _, ok := Apps[pid][aid]; !ok {
		return Enums.NOT_FOUND_APP
	}

	return nil
}

func InitAppChan(pid string) bool {
	if _, ok := AppChans[pid]; !ok {
		chanTmp := make(chan int)
		lockTmp := new(sync.Mutex)
		var appChan = AppChan{chanTmp, lockTmp}
		AppChans[pid] = &appChan

		return true
	}

	return false
}

func PutApp(pid string, aid string, name string, descr string) bool {
	if _, ok := Apps[pid]; !ok {
		Apps[pid] = make(map[string]*App)
	}

	var have bool
	if _, ok := Apps[pid][aid]; !ok {
		Apps[pid][aid] = new(App)
		Apps[pid][aid].CreateTime = time.Now().Unix()
		have = true
	}

	if !have {
		if Apps[pid][aid].Name == name && Apps[pid][aid].Description == descr && Helpers.FileExists(Enums.GetDBAppsFile(pid)) {
			return true
		}
	}

	AppChans[pid].Lock.Lock()
	Apps[pid][aid].Name = name
	Apps[pid][aid].Description = descr
	Apps[pid][aid].UpdateTime = time.Now().Unix()
	AppChans[pid].Lock.Unlock()

	AppChans[pid].Channel <- 1

	return true
}

func DeleteApp(pid string, aid string) bool {
	AppChans[pid].Lock.Lock()
	delete(Apps[pid], aid)
	AppChans[pid].Lock.Unlock()

	AppChans[pid].Channel <- 1

	return true
}
