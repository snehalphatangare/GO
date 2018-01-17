package main
import (
	"fmt";"sync"
)

type SingletonDoubleCheckLock struct{}

var single *SingletonDoubleCheckLock
var sync_var sync.Mutex

func getInstance() *SingletonDoubleCheckLock{
	sync_var.Lock()
	defer sync_var.Unlock()
	if    single == nil {
		single = &SingletonDoubleCheckLock{}
	}
	return single 
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
