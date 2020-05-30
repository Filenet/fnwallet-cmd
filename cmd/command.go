package cmd

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

var Chanl = make(chan int)

func Command() {
	os.Stdout.WriteString("filenet>")
	for {
		input := bufio.NewReader(os.Stdin)
		args, _, _ := input.ReadLine()
		command := string(args)
		comm := strings.Fields(command)
		if len(comm) == 0 {
			continue
		}
		switch comm[0] {
		case "ipfs":
			IPFSserver(comm[1:])
			FilenetHeader(nil)
		case "filenet":
			FilenetServer(comm[1:])
			FilenetHeader(nil)
		case "bye":
			FilenetExit(nil)
		case "exit":
			FilenetExit(nil)
		case "quit":
			FilenetExit(nil)
		default:
			cmd := exec.Command("system", comm...)
			cmd.Run()
			FilenetHeader(nil)
		}
	}
}

