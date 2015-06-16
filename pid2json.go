package pid2json

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-ps"
)

type Pidstream struct {
	Pid   int
	Pname string
}

// print array of PIDs
func tap() {
	arrayofPIDs, _ := ps.Processes()

	for _, PID := range arrayofPIDs {
		fmt.Print(PID)
	}
}

func nozzle() {}
