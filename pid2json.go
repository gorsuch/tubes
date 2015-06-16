package pid2json

import (
	//	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-ps"
)

// print array of PIDs
func tap() {
	arrayofPIDs, _ := ps.Processes()

	for _, PID := range arrayofPIDs {
		fmt.Print(PID)
	}
}

func Nozzle() {}
