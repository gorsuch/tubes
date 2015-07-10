package main

//package tubes

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mitchellh/go-ps"
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
		fmt.Println(reflect.TypeOf(m))
		b, _ := json.Marshal(m)
		fmt.Println(string(b))
		fmt.Println(reflect.TypeOf(b))
		fmt.Println(reflect.TypeOf(b[0]))
	}
}

func Nozzle() {

}

func main() {
	tap()
}
