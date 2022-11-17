package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() *single {
	// fmt.Println(singleInstance)
	lock.Lock()

	defer lock.Unlock()
	if singleInstance == nil {

		fmt.Println("Creating single instance now.")
		singleInstance = &single{}
		return singleInstance
	}

	fmt.Println("Single instance already created.")
	return singleInstance
}
