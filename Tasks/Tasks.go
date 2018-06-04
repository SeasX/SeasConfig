package Tasks

import (
	"github.com/SeasX/SeasConfig/Models"
)

func init() {
	LoadProducts(&Models.Products)
	LoadApps()
	LoadConfigs()

	WatchSaveProducts()
}

func WatchSaveProducts() {
	go func() {
		for {
			select {
			case <-Models.ProductsChan:
				go func() {
					SaveProducts(Models.GetAllProducts())
					return
				}()
			}
		}
	}()
}

func WatchSaveApps(pid string) {
	go func(pid string) {
		for {
			select {
			case <-Models.AppChans[pid].Channel:
				go func(pid string) {
					SaveApps(pid, Models.GetAllApps(pid))
					return
				}(pid)
			}
		}
	}(pid)
}

func WatchSaveConfigs(pid string, aid string) {
	go func(pid string, aid string) {
		for {
			select {
			case <-Models.ConfigChans[pid+aid].Channel:
				go func(pid string, aid string) {
					SaveConfigs(pid, aid, Models.GetAppConfigs(pid, aid))
					return
				}(pid, aid)
			}
		}
	}(pid, aid)
}
