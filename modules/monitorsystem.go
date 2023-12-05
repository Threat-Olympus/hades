// Real-time monitoring of system events
package modules

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Monitor Nework Events
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

// Monitor Windows Log Events
func LogEvents() {
	cmd := exec.Command("wevtutil", "qe", "System", "/f:text")

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

		if strings.Contains(line, "Error") {
			fmt.Println("Potential threat detected:", line)
		}
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for Cmd", err)
		return
	}
}

// Monitor CPU Usage
func MonitorCPU() {
	for {
		// Print the current CPU usage
		fmt.Printf("Current CPU usage: %f%%\n", float64(runtime.NumCPU())/float64(runtime.NumGoroutine())*100)
		time.Sleep(1 * time.Second)
	}
}

// Monitor Memory Usage
func MonitorMemory() {
	for {
		// Print the current memory usage
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Current memory usage: %v bytes\n", m.Alloc)
		time.Sleep(1 * time.Second)
	}
}

// Monitor File System
func MonitorFileSystem(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
