package vlcremote

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// template := `/Applications/VLC.app/Contents/MacOS/VLC --intf http --http-host 127.0.0.1 --http-port %d --http-password %s`
func OpenVlc(port int, password string) (*exec.Cmd, error) {
	path := "/Applications/VLC.app/Contents/MacOS/VLC"
	vlcCmd := exec.Command(
		path,
		"--intf", "macosx",
		"--extraintf", "http",
		"--http-host", "127.0.0.1",
		"--http-port", fmt.Sprintf("%d", port),
		"--http-password", password,
	)

	/*
		vlcCmd.Stderr = os.Stdout
	*/
	vlcCmd.Stdout = os.Stdout

	errPipe, err := vlcCmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(errPipe)

	if err := vlcCmd.Start(); err != nil {
		return nil, err
	}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Read line:", line)
		if strings.Contains(line, "[http] lua interface: Lua HTTP interface") {
			break
		}
	}

	vlcCmd.Stderr = os.Stderr

	return vlcCmd, nil
}
