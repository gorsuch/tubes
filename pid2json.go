package main

//package tubes

import (
	"encoding/json"
	//"fmt"
	"net/http"
	//	"reflect"

	"github.com/mitchellh/go-ps"
)

type PsData struct {
	Pid  int    `json:"pid"`
	PPid int    `json:"ppid"`
	Exec string `json:"exec"`
}

var PsMarshaled []byte
var PsCollection []PsData

// take []ps.Process and format it into a slice of JSON values
func tap() {

	PsSnapshot, _ := ps.Processes()

	// PsSnapshot is now type []ps.Process
	// Convert to slice of JSON formatted es
	// using our PsData struct

	for _, e := range PsSnapshot {
		data := PsData{
			Pid:  e.Pid(),
			PPid: e.PPid(),
			Exec: e.Executable(),
		}
		PsCollection = append(PsCollection, data)
	}
}

// Nozzle sprays out some JSON
func Nozzle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(PsCollection)
}

func main() {
	// populate PsCollection
	tap()

	// let's run Nozzle() when folks hit webserver root
	http.HandleFunc("/", Nozzle)

	// start webserver
	http.ListenAndServe(":8080", nil)
}
