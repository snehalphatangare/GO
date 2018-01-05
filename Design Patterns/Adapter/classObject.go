package main

import "fmt"

//Target interface
type Target interface {
	sayHello()
}

type AdapteeParent struct {}

func (AdapteeParent) printMessage(msg string)  {
	fmt.Println(msg)
}

type Adapter struct {
	AdapteeParent 
}

func (Adapter) sayHello() {
	a := Adapter{AdapteeParent{}}
	a.printMessage("Hello")
	
}

func runAdapter(){
		a1 := Adapter{AdapteeParent{}}
		a1.sayHello()
}

func main(){
	runAdapter()
}