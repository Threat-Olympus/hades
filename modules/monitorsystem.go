// Real-time monitoring of system events
package modules

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func NetworkEvents() {
	cmd := exec.Command("netstat", "-a")

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
		}
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for Cmd", err)
		return
	}
}
