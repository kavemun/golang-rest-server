package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func StartJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Starting Job...\n")

	cmd := exec.Command("go", "version")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	// Execute command
	printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	printError(err)
	printOutput(cmdOutput.Bytes())

	todos := Todos{
		Todo{Name: "Start Job"},
	}

	json.NewEncoder(w).Encode(todos)

}

func PauseJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pausing Job...\n")

	cmd := exec.Command("go", "version")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	// Execute command
	printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	printError(err)
	printOutput(cmdOutput.Bytes())

	todos := Todos{
		Todo{Name: "Pause Job"},
	}

	json.NewEncoder(w).Encode(todos)

}

func TerminateJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Terminating Job...\n")

	cmd := exec.Command("go", "version")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	// Execute command
	printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	printError(err)
	printOutput(cmdOutput.Bytes())

	todos := Todos{
		Todo{Name: "Terminate presentation"},
	}

	json.NewEncoder(w).Encode(todos)

}
