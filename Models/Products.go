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
	Products    map[string]*Product
	ProductLock map[string]*sync.Mutex
)

func init() {
	Products = make(map[string]*Product)
	ProductLock = make(map[string]*sync.Mutex)
}

func GetAllProducts() (maps map[string]*Product) {
	return Products
}

func GetProduct(pid string) (pro *Product) {
	return Products[pid]
}

func PutProduct(pid string, name string, descr string) bool {
	if _, ok := ProductLock[pid]; !ok {
		ProductLock[pid] = new(sync.Mutex)
	}
	ProductLock[pid].Lock()
	defer ProductLock[pid].Unlock()

	if _, ok := Products[pid]; !ok {
		Products[pid] = new(Product)
		Products[pid].CreateTime = time.Now().Unix()
	}

	Products[pid].Name = name
	Products[pid].Description = descr
	Products[pid].UpdateTime = time.Now().Unix()

	return true
}

func DeleteProduct(pid string) bool {
	ProductLock[pid].Lock()
	defer ProductLock[pid].Unlock()

	delete(Products, pid)

	return true
}
