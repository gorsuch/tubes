package main

//package tubes

import (
	//"encoding/json"
	"fmt"
	"github.com/mitchellh/go-ps"
	"reflect"
	"strconv"
	//"strings"
)

// print array of PIDs
func tap() {

	PIDs, _ := ps.Processes()
	//fmt.Println(reflect.TypeOf(PIDs))
	//fmt.Println(PIDs)

	aPID := PIDs[0]
	fmt.Println(reflect.TypeOf(&aPID))
	fmt.Println(aPID.Pid())
	b := aPID.Pid()
	fmt.Println(reflect.TypeOf(b))

	b_string := strconv.Itoa(b)
	fmt.Println(b_string)
	fmt.Println(reflect.TypeOf(b_string))

	//	for _, PID := range PIDs {
	//		fmt.Println(PID)
	//		fmt.Println(reflect.TypeOf(PID))
	//	}

}

func Nozzle() {

}

func main() {
	tap()
}
