package main

import (
	"driver"
	"time"
	"fmt"
)

type order struct {
	floor int
	orderType int // 0 for up, 1 for down and 2 for command
	//elevatorId int //Which elevator the order is from
}




func CheckUpButtons(ch chan order) {	
	for i := 0; i < 3; i++ {
		if driver.GetButtonSignal(0,i) == 1 {
			driver.SetButtonLampOn(0,i)
			newOrder := order{floor:i, orderType:0}
			ch <- newOrder
		}
	}
}

func CheckDownButtons(ch chan order) {	
	for i := 1; i < 4; i++ {
		if driver.GetButtonSignal(1,i) == 1 {
			driver.SetButtonLampOn(1,i)
			newOrder := order{floor:i, orderType:1}
			ch <- newOrder
		}
	}
}

func CheckCommandButtons(ch chan order) {
	for j := 0; j < 4; j++ {
		if driver.GetButtonSignal(2,j) == 1 {
			driver.SetButtonLampOn(2,j)
			newOrder := order{floor:j, orderType:2}
			ch <- newOrder
		}
	}
}

func ButtonPoller(ch chan order) {
	for {
		CheckDownButtons(ch)
		CheckUpButtons(ch)
		CheckCommandButtons(ch)
		time.Sleep(10)
	}
}

func FloorPoller(ch chan int) {
	for {
		currentFloor := driver.GetFloor()
		switch  currentFloor {
		case 0:
			driver.SetFloorLamp(0)
			ch <- 0
		case 1:
			driver.SetFloorLamp(1)
			ch <- 1
		case 2:
			driver.SetFloorLamp(2)
			ch <- 2
		case 3:
			driver.SetFloorLamp(3)
			ch <- 3
		}		
		time.Sleep(10)
	}		
}	
	
func testRun(ch chan int) {
	driver.SetDirection(-1)
	for {
		floor := <- ch
		if floor == 3 {
			driver.SetDirection(-1)
		} else if floor == 0 {		
			driver.SetDirection(1)
		}
		time.Sleep(10)
	}
}

func PrintOrders(ch chan order) {
	for {
		newOrder := <- ch 
		fmt.Printf("Floor: %v \n", newOrder.floor)	
		fmt.Printf("Button:%v \n", newOrder.orderType)
		time.Sleep(10)
	}
}


func main () {
	driver.Init()
	/*messages := make(chan order)
	go ButtonPoller(messages)
	go PrintOrders(messages)
	*/
	floor := make(chan int)
	go FloorPoller(floor)
	go testRun(floor)
	neverReturn := make (chan int)
	<-neverReturn
}
				
		
	
