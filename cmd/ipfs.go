package cmd

import (
	"os"
	"os/exec"
	"time"
)

func IPFSserver(args []string) {
	go func() {
		cmd := exec.Command("ipfs")
		for _, a := range args {
			cmd.Args = append(cmd.Args, a)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr=os.Stderr
		err := cmd.Start()
		cmd.Wait()
		if err != nil {
			res:="invalid option\nTry executing \"ipfs --help\" for more information."
			cmd.Stderr.Write([]byte(res))
		}
		time.Sleep(time.Second)
	}()
}