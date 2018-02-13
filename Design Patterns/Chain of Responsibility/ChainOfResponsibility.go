package main

import "fmt"

type Handler interface{
	handleRequest(request string)
	SetSuccessor(next Handler)
}

type ConcreteHandler1 struct{
	Successor Handler
}

//ConcreteHandler1 implelmenting Handler interface
func (concrete1 ConcreteHandler1) handleRequest(request string) {
 	fmt.Printf("R1 got the request...") 
	if ( request == "R1" ){
            fmt.Println(" => This one is mine!")
        }else{
	     fmt.Println(" => This one is not mine!")
            if ( concrete1.Successor != nil ){
				concrete1.Successor.handleRequest(request);
			}	
                
        }
}

func (concrete1 *ConcreteHandler1) SetSuccessor(next Handler) {
        concrete1.Successor = next 
	}

//ConcreteHandler2 implelmenting Handler interface
type ConcreteHandler2 struct{
	Successor Handler
}

func (concrete2 ConcreteHandler2) handleRequest(request string) {
 	fmt.Println("R2 got the request...") 
	if ( request == "R2" ){
            fmt.Println(" => This one is mine!")
        }else{
	     fmt.Println(" => This one is not mine!")
         if ( concrete2.Successor != nil ){
				concrete2.Successor.handleRequest(request);
		}	
                
        }
}

func (concrete2 *ConcreteHandler2) SetSuccessor(next Handler) {
        concrete2.Successor = next 
	}


//ConcreteHandler3 implelmenting Handler interface
type ConcreteHandler3 struct{
	Successor Handler
}

func (concrete3 ConcreteHandler3) handleRequest(request string) {
 	fmt.Println("R3 got the request...") 
	if ( request == "R3" ){
            fmt.Println(" => This one is mine!")
        }else{
	     fmt.Println(" => This one is not mine!")
         if ( concrete3.Successor != nil ){
			concrete3.Successor.handleRequest(request);
		}	
                
    }
}

func (concrete3 *ConcreteHandler3) SetSuccessor(next Handler) {
        concrete3.Successor = next 
	}


func runChainOfResponsibility(){
	ch1 := new(ConcreteHandler1)
	ch2 := new(ConcreteHandler2)
	ch3 := new(ConcreteHandler3)
	
	ch1.SetSuccessor(ch2)
	ch2.SetSuccessor(ch3)
	
	fmt.Println("Sending R2...")
	ch1.handleRequest("R2")
	fmt.Println("Sending R3...")
	ch1.handleRequest("R3")
	fmt.Println("Sending R1...")
	ch1.handleRequest("R1")
	fmt.Println("Sending RX...")
	ch1.handleRequest("RX")
}

func main(){
	runChainOfResponsibility()
}
