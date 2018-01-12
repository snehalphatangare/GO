package main

import (
	"fmt";
	"sync"
)
type singleton struct{
}

var instance *singleton

var once sync.Once

func getInstance() *singleton{
	once.Do(func() {
	        if (instance == nil){
			instance =&singleton{}
		}
	})
	return instance 
}

func runSingleton(){
	instance1 := getInstance();
	instance2 := getInstance();
	if instance1 == instance2 {
		fmt.Println("Both the instances are same")
	}
} 

func main(){
	runSingleton()
}