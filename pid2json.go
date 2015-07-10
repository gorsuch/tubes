package main

//package tubes

import (
	"encoding/json"
	"fmt"

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

// print array of Pidsnapshot
func tap() {

	Pidsnapshot, _ := ps.Processes()

	// Pidsnapshot is now type []ps.Process
	// Convert to JSON

	for _, element := range Pidsnapshot {
		m := PidValues{Pid: element.Pid(), PPid: element.PPid(), Exec: element.Executable()}
		b, _ = json.Marshal(m)
	}

	fmt.Println(string(b))
}

func Nozzle() {

}

func main() {
	tap()
}
