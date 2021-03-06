package main

import "fmt"
//Prathmesh
var successorForOne handler
var successorForTwo handler
var successorForThree handler

type handler interface{
	handleRequest(request string)
	setSuccessor(next handler)
}

type controller1 struct {
	R1 string
}

type controller2 struct {
	R2 string
}

type controller3 struct {
	R3 string
}

func (c1 controller1) handleRequest(request string) {
fmt.Println("R1 got the request...")
	if request == "R1"{
	fmt.Println("controller1 => This one is mine!")
	}else {
		if successorForOne != nil {
			successorForOne.handleRequest(request)
		}
	}
	}

func (c1 controller1)setSuccessor(next handler){

		successorForOne = next
}

func (c2 controller2) handleRequest(request string) {
	fmt.Println("R2 got the request...")
	if request == "R2"{
		fmt.Println("controller2 => This one is mine!")
	} else {
		if successorForTwo != nil {
			successorForTwo.handleRequest(request)
		}
	}
}
func (c2 controller2)setSuccessor(next handler){

	successorForTwo  = next

}

func (c3 controller3) handleRequest(request string){
	fmt.Println("R3 got the request...")
	if request == "R3"{
		fmt.Println("controller3 => This one is mine!")
	} else {
		if successorForThree != nil {
		successorForThree.handleRequest(request)
		}
	}
}
func (c3 controller3)setSuccessor(next handler){

	successorForThree  = next

}

func runTest(h1 handler, h2 handler, h3 handler){
h1.setSuccessor(h2)
h2.setSuccessor(h3)
fmt.Println("Sending R2...")
h1.handleRequest("R2")

fmt.Println("Sending R3...")
h1.handleRequest("R3")

fmt.Println("Sending R1...")
h1.handleRequest("R1")

//fmt.Println("Sending RX...")
h1.handleRequest("RX")

}

func main(){
c1 := controller1{"R1"}
c2 := controller2{"R2"}
c3 := controller3{"R3"}
fmt.Print(successorForThree)
runTest(c1, c2, c3)
}
