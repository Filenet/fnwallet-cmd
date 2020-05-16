package cmd

import (
  "go.uber.org/zap/buffer"
  "os"
  "os/exec"
  "time"
)

func IPFSserver(args []string) {
	cmd := exec.Command("ipfs")
	for _, a := range args {
		cmd.Args = append(cmd.Args, a)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if cmd.Args[1] == "daemon" {
      var i = make(chan int)
		go func() {
		    i<-1
			err := cmd.Run()
			if err != nil {
				res := "invalid option\nTry executing \"ipfs --help\" for more information."
				cmd.Stderr.Write([]byte(res))
			}
		}()
      <-i
      time.Sleep(time.Second)
	} else {
		err := cmd.Start()
		if err != nil {
			res := "invalid option\nTry executing \"ipfs --help\" for more information."
			cmd.Stderr.Write([]byte(res))
		}
		cmd.Wait()
	}
}

func IPFSstart(cmd *exec.Cmd, stdOut, stdErr *buffer.Buffer) {
	cmd.Stdout = stdOut
	cmd.Stderr = stdErr
	err := cmd.Run()
	if err != nil {
		res := "invalid option\nTry executing \"ipfs --help\" for more information."
		cmd.Stderr.Write([]byte(res))
	}
	time.Sleep(time.Second * 2)
}
