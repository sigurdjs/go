package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go
/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "elev.h"
*/
import "C"

func Init(){
	C.elev_init()
} 

func SetDirection(direction int) {
	C.elev_set_motor_direction(C.int(direction))
}

func SetDoorOpen(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}

func SetStopLamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func SetFloorLamp(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func GetFloor() int {
	return int(C.elev_get_floor_sensor_signal())
}

func GetObstruction() int {
	return int(C.elev_get_obstruction_signal())
}

func GetStopSignal() int {
	return int(C.elev_get_stop_signal())
}


//Buttontype = 0 for up, 1 for down and 2 for command
func SetButtonLampOn(buttonType int, floor int) { 
	C.elev_set_button_lamp(C.int(buttonType),C.int(floor),1)
}

//Buttontype = 0 for up, 1 for down and 2 for command
func SetButtonLampOff(buttonType int ,floor int) { 
	C.elev_set_button_lamp(C.int(buttonType),C.int(floor),0)
}


//Buttontype = 0 for up, 1 for down and 2 for command
func GetButtonSignal(buttonType int, floor int) int {
	return int(C.elev_get_button_signal(C.int(buttonType),C.int(floor)))
}

	














