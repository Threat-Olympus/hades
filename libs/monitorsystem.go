package libs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func MONITORSYSTEMEVENT() {
	cmd := exec.Command("netstat", "-nt")

	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Cmd", err)
		return
	}

	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "LISTEN") {
			fmt.Println("Potential threat detected:", line)
		} else {
			fmt.Println("No threat detected")
			os.Exit(0)
		}
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for Cmd", err)
		return
	}
}
