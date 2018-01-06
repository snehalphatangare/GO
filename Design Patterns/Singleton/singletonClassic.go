package main

import "fmt"

type singleton struct{}

var singleInstance *singleton

func getInstance() *singleton{
    if (singleInstance == nil){
		singleInstance =&singleton{}
	}
	return singleInstance 
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
