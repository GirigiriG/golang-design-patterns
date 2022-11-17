package main

import (
	"singleton-pattern/singleton"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go singleton.GetInstance()
	}
	time.Sleep(time.Microsecond * 500)
}
