package main

//package tubes

import (
	"encoding/json"
	"fmt"
	//	"reflect"

	"github.com/mitchellh/go-ps"
	//"strconv"
	//"strings"
)

type PidValues struct {
	Pid  int
	PPid int
	Exec string
}

var b []byte
var a []string

// print array of Pidsnapshot
func tap() {

	Pidsnapshot, _ := ps.Processes()

	// Pidsnapshot is now type []ps.Process
	// Convert to slice of JSON formatted elements

	for _, element := range Pidsnapshot {
		m := PidValues{Pid: element.Pid(), PPid: element.PPid(), Exec: element.Executable()}
		b, _ = json.Marshal(m)
		a = append(a, string(b))
	}
	fmt.Println(a)
}

func Nozzle() {

}

func main() {
	tap()
}
