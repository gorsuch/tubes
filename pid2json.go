package main

//package tubes

import (
	"encoding/json"
	"fmt"
	"net/http"
	//	"reflect"

	"github.com/mitchellh/go-ps"
	//"strconv"
	//"strings"
)

type PsValues struct {
	Pid  int
	PPid int
	Exec string
}

var PsMarshaled []byte
var PsJson []string

// take []ps.Process and format it into a slice of JSON values
func tap() {

	PsSnapshot, _ := ps.Processes()

	// PsSnapshot is now type []ps.Process
	// Convert to slice of JSON formatted elements

	for _, element := range PsSnapshot {
		PsMap := PsValues{Pid: element.Pid(), PPid: element.PPid(), Exec: element.Executable()}
		PsMarshaled, _ = json.Marshal(PsMap)
		PsJson = append(PsJson, string(PsMarshaled))
	}
	//fmt.Println(PsJson)
}

func Nozzle(w http.ResponseWriter, r *http.Request) {
	for _, s := range PsJson {
		fmt.Fprintf(w, string(s))
	}
}

func main() {
	// populate PsJson
	tap()

	// let's run Nozzle() when folks hit webserver
	http.HandleFunc("/", Nozzle)

	// start webserver
	http.ListenAndServe(":8080", nil)
}
