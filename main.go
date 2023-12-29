package main

import (
	"flag"
	"fmt"
	"hades/modules"
	"os"
	"time"
)

var (
	help   *bool
	net    *bool
	event  *bool
	cpu    *bool
	mem    *bool
	fsm    *bool
	path   *string
	threat *bool
)

func init() {
	help = flag.Bool("help", false, "Show help")
	net = flag.Bool("net", false, "Monitor network events")
	event = flag.Bool("event", false, "Monitor Windows log events")
	cpu = flag.Bool("cpu", false, "Monitor CPU usage")
	mem = flag.Bool("mem", false, "Monitor memory usage")
	fsm = flag.Bool("fsm", false, "Monitor file system events")
	path = flag.String("path", ".", "path to file monitor")
	threat = flag.Bool("threat", false, "Detect known threat indicators")
}

func main() {

	messageChan := make(chan string)

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *net {
		messageChan <- modules.NetworkEvents()
		os.Exit(0)
	}

	if *event {
		messageChan <- modules.LogEvents()
		os.Exit(0)
	}

	if *cpu {
		messageChan <- modules.MonitorCPU()
	}

	if *mem {
		messageChan <- modules.MonitorMemory()
	}

	if *fsm {
		if *path != "." {
			messageChan <- modules.MonitorFileSystem(*path)
		} else {
			fmt.Println("[-]: Path is not set")
		}
	}

	if *threat {
		fmt.Println("[+] Checking for DLL injection...")
		messageChan <- modules.CheckDllInjection()
		time.Sleep(2 * time.Second)
		fmt.Println("[+] Checking for process injection...")
		messageChan <- modules.CheckProcessInjection()
	}

	message := <-messageChan
	fmt.Println(message)
}
