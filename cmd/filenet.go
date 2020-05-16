package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fnv3/test/filenetipfs"
	"os"
	"path/filepath"
	"sync"
)

var (
	Filenet = make(map[string]func(args []string))
)

func init() {
	Filenet["help"] = FilenetHelp
	Filenet["add"] = FilenetAdd
	Filenet["get"] = FilenetGet
	Filenet["set"] = FilenetSet
}

func FilenetServer(args []string) {
	if _, ok := Filenet[args[0]]; !ok {
		res := fmt.Sprintf("%s\t%s", INVALIDOPTION, args[0])
		os.Stderr.Write([]byte(res))
		return
	}
	Filenet[args[0]](args[1:])
}

func FilenetCommandMap(command string, function func(args []string)) {
	if _, ok := Filenet[command]; !ok {
		Filenet[command] = function
	}
}

func FilenetHelp(args []string) {
	Hint := fmt.Sprintf("%s\n%s\t%s\n", "USAGE", "filenet", "- filenet config")
	fmt.Println(Hint)
}

func FilenetAdd(args []string) {
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	var wait sync.WaitGroup
	wait.Add(1)
	_, fileName := filepath.Split(args[0])
	go func() {
		res, err := filenetipfs.SaveFile(file, fileName)
		if err != nil {
			fmt.Println(err)
		}
		r, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Please do not close")
		os.Stdout.Write(r)
		wait.Done()
	}()
	wait.Wait()
	fmt.Println("success")
	return

}

func FilenetExit(args []string) {
	Chanl <- 1
	return
}

func FilenetHeader(args []string) {
	out := bufio.NewReader(os.Stdout)
	b, err := out.ReadByte()
	fmt.Print(b, err)
	if err != nil && b != 0 {
		os.Stdout.WriteString("filenet>")
	} else {
		os.Stdout.WriteString("\nfilenet>")
	}
}

func FilenetGet(args []string) {
	config,err:=NewConfig()
	if err!=nil{
		os.Stderr.Write([]byte(err.Error()))
		return
	}
	str,err:=config.String()
	if err!=nil{
		os.Stderr.Write([]byte(err.Error()))
		return
	}
	os.Stdout.Write([]byte(str))
	return
}

func FilenetSet(args []string) {
	config, err := set(args[1])
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		return
	}
	b, err := json.Marshal(config)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		return
	}
	os.Stdout.Write(b)
	return
}


