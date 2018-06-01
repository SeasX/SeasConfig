package Models

import (
	"sync"
	"time"
)

type Product struct {
	Name        string
	Description string
	CreateTime  int64
	UpdateTime  int64
}

var (
	Products     map[string]*Product
	ProductLock  map[string]*sync.Mutex
	ProductsChan chan int
)

func init() {
	Products = make(map[string]*Product)
	ProductLock = make(map[string]*sync.Mutex)
	ProductsChan = make(chan int)
}

func GetAllProducts() (maps map[string]*Product) {
	return Products
}

func ExistsProducts() (have bool) {
	if len(Products) < 1 {
		return false
	}

	return true
}

func GetProduct(pid string) (pro *Product, have bool) {
	if _, ok := Products[pid]; !ok {
		return nil, false
	}
	return Products[pid], true
}

func ExistsProduct(pid string) bool {
	if _, ok := Products[pid]; !ok {
		return false
	}

	return true
}

func PutProduct(pid string, name string, descr string) bool {
	if _, ok := ProductLock[pid]; !ok {
		ProductLock[pid] = new(sync.Mutex)
	}
	ProductLock[pid].Lock()
	defer ProductLock[pid].Unlock()

	var have bool
	if _, ok := Products[pid]; !ok {
		Products[pid] = new(Product)
		Products[pid].CreateTime = time.Now().Unix()
		have = true
	}

	if !have {
		if Products[pid].Name == name && Products[pid].Description == descr {
			return true
		}
	}

	Products[pid].Name = name
	Products[pid].Description = descr
	Products[pid].UpdateTime = time.Now().Unix()

	ProductsChan <- 1

	return true
}

func DeleteProduct(pid string) bool {
	ProductLock[pid].Lock()
	defer ProductLock[pid].Unlock()

	delete(Products, pid)

	ProductsChan <- 1

	return true
}
