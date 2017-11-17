package main

import "fmt"

type Singleton struct {
}

var singleInstance *Singleton
var instance *Singleton

func getInstance() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

func main() {

	singleInstance = getInstance()
	fmt.Println(&singleInstance)

	singleInstance = getInstance()
	fmt.Println(&singleInstance)
}
