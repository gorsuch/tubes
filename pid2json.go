package main

//package tubes

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-ps"
	"reflect"
	//"strconv"
	//"strings"
)

type PidValues struct {
	Pid  int
	PPid int
	Exec string
}

// print array of PIDs
func tap() {

	PIDs, _ := ps.Processes()

	// PIDs is now type []ps.Process
	// Convert to JSON
	for _, element := range PIDs {
		m := PidValues{Pid: element.Pid(), PPid: element.PPid(), Exec: element.Executable()}
		fmt.Println(m)
		b, _ := json.Marshal(m)
		fmt.Println(reflect.TypeOf(b))
		//fmt.Println(b)
	}
}

func Nozzle() {

}

func main() {
	tap()
}
