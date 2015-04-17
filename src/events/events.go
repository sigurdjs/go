package main

import (
	"driver"
	"time"
	"fmt"
	"Network"
	"udp"
)

var LightArray [3][4]int //row 0 for up, row 1 for down, row 2 for inside

type order struct {
	floor int
	orderType int // 0 for up, 1 for down and 2 for command
	//elevatorId int //Which elevator the order is from
}



func CheckUpButtons(send_ch chan udp.Udp_message) {	
	for i := 0; i < 3; i++ {
		if driver.GetButtonSignal(0,i) == 1 && LightArray[0][i] == 0{
			LightArray[0][i] = 1
			driver.SetButtonLampOn(0,i)
			Network.SendNewOrderMessage(send_ch,1,0,i)		
		}
	}
}

func CheckDownButtons(send_ch chan udp.Udp_message) {	
	for i := 1; i < 4; i++ {
		if driver.GetButtonSignal(1,i) == 1 && LightArray[1][i] == 0{
			LightArray[1][i] = 1
			driver.SetButtonLampOn(1,i)			
			Network.SendNewOrderMessage(send_ch,1,1,i)
		}
	}
}

func CheckCommandButtons(send_ch chan udp.Udp_message) {
	for i := 0; i < 4; i++ {
		if driver.GetButtonSignal(2,i) == 1 && LightArray[2][i] == 0 {
			LightArray[2][i] = 1			
			driver.SetButtonLampOn(2,i)
			Network.SendNewOrderMessage(send_ch,1,2,i)
		}
	}
}

func ButtonPoller(send_ch chan udp.Udp_message) {
	for {
		CheckDownButtons(send_ch)
		CheckUpButtons(send_ch)
		CheckCommandButtons(send_ch)
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
	send_ch := make (chan udp.Udp_message)
	receive_ch := make (chan udp.Udp_message)
	err := udp.Udp_init(20019, 20019, 100, send_ch, receive_ch)	
	go ButtonPoller(send_ch)
	/*floor := make(chan int)
	go FloorPoller(floor)
	go testRun(floor)*/
	if (err != nil){
		fmt.Print("main done. err = %s \n", err)
	}
	neverReturn := make (chan int)
	<-neverReturn
}
				
		
	
