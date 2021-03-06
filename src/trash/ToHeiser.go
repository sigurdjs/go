package main 


import (
	"fmt"
	"udp"
	"time"
	"driver"
	"strings"
)

var message string
var startLift bool

func print_udp_message(msg udp.Udp_message){ 	
	fmt.Printf("msg:  \n \t raddr = %s \n \t data = %s \n \t length = %v \n", msg.Raddr, msg.Data, msg.Length)
}

func waitAndStart() {
	for {
		if driver.GetButtonSignal(2,0) == 1 {
			message = "Start Lift"
		} else {			
			message = "stopp lift"
		}
		time.Sleep(100)	
	}
}


func node (send_ch, receive_ch chan udp.Udp_message){
	message = "stopp lift"
	for {
		time.Sleep(1*time.Second)//78.91.45.202:20002
		//problemet er å skrive til meg selv på min "lokale" port
		snd_msg := udp.Udp_message{Raddr:"129.241.187.150:20003", Data:message, Length:10}
		fmt.Printf("Sending------\n")
		send_ch <- snd_msg
		print_udp_message(snd_msg)
		fmt.Printf("Receiving----\n")
		rcv_msg:= <- receive_ch
		print_udp_message(rcv_msg)
		if strings.Contains(rcv_msg.Data, "Start Lift") {
			startLift = true
		}		
	}
}

func runLift() {
	for {
		if startLift == true {
			var floor int
			driver.SetDirection(1)
			for {
				floor = driver.GetFloor() 
				if floor != -1 {
					driver.SetFloorLamp(floor)		
				}
				if floor == 3 {
					driver.SetButtonLampOn(2,1)
					driver.SetDirection(-1)
				} else if floor == 0 {
					driver.SetButtonLampOff(2,1)
					driver.SetDirection(1)
				} else if driver.GetStopSignal() == 1 {
					driver.SetDirection(0)
					startLift = false
					message = "stopp lift"
					break
				}
				time.Sleep(100)
			}	
		}		
		time.Sleep(100)
	}
}	
 


func main (){	
	driver.Init()
	send_ch := make (chan udp.Udp_message)
	receive_ch := make (chan udp.Udp_message)
	err := udp.Udp_init(20002, 20003, 1024, send_ch, receive_ch)	
	go node(send_ch, receive_ch)
	go waitAndStart()
	go runLift()
	

	if (err != nil){
		fmt.Print("main done. err = %s \n", err)
	}
		neverReturn := make (chan int)
	<-neverReturn

}
