package pid2jsontube

import (
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

func nozzle() {}
