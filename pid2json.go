package main

//package tubes

import (
	//"encoding/json"
	"fmt"
	"github.com/mitchellh/go-ps"
	"reflect"
	//"strings"
)

// print array of PIDs
func tap() {

	PIDs, _ := ps.Processes()

	for _, PID := range PIDs {
		fmt.Println(PID)
		fmt.Println(reflect.TypeOf(PID))
	}

}

func Nozzle() {

}

func main() {
	tap()
}
