package Tasks

import (
	"github.com/SeasX/SeasConfig/Models"
)

func init() {
	LoadProducts(&Models.Products)
	LoadApps()

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
