package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	taskL := exec.Command("tasklist")

	var out bytes.Buffer
	taskL.Stdout = &out
	if err := taskL.Run(); err != nil {
		log.Fatal(err)
	}
	var processId string
	output := out.String()

	lines := strings.Split(output, "\n")
	fmt.Println(len(lines))

	for _, line := range lines {
		if strings.Contains(line, "vlc.exe") {
			fields := strings.Fields(line)
			fmt.Println(fields)

			if len(fields) >= 2 {
				processId = fields[1]
				fmt.Println(processId)
			}
		}
	}

	pid, err := strconv.Atoi(processId)
	if err != nil {
		log.Fatal(err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}

	err = proc.Kill()
	if err != nil {
		log.Fatalf("Error killing process: %v", err)
	}
}
