package modules

import (
	"bufio"
	"fmt"
	"os/exec"
)

func CheckDllInjection() {
	script := `
    $Path = 'C:\Windows\System32'
    $KnownDLLs = Get-Content 'HKLM:\System\CurrentControlSet\Control\Session Manager\KnownDLLs'
    $KnownDLLs | ForEach-Object {
        $FullPath = Join-Path -Path $Path -ChildPath $_.PSChildName
        if (!(Test-Path -Path $FullPath)) {
            Write-Output ("Potential DLL hijacking detected: " + $FullPath)
        }
    }
    `

	cmd := exec.Command("powershell", "-Command", script)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}

func CheckProcessInjection() {
	script := `
	$Processes = Get-Process
	$Processes | ForEach-Object {
		$ProcessPath = $_.Path
		if ($ProcessPath -eq $null) {
			Write-Output ("Potential process injection detected: " + $_.Name)
		}
	}
	`

	cmd := exec.Command("powershell", "-Command", script)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}
