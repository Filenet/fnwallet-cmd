package cmd

import (
	"fmt"
	"fnv3/test/filenetipfs"
	"os"
	"sync"
)

var (
	Filenet = make(map[string]func(args []string))
)

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
	fmt.Println(file.Name())
	wait := new(sync.WaitGroup)
	wait.Add(1)
	go func() {
		err = filenetipfs.SaveFile(file, "11")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Please do not close")
		wait.Done()
	}()
	fmt.Println("success")
	return

}

func FilenetGet(args []string) {
	Hint := fmt.Sprintf("%s\n%s\t%s\n", "USAGE", "filenet", "- filenet config")
	fmt.Println(Hint)
}

//func FilnetAPIPort(args []string) {
//	port := args[0]
//	err := Config.Set("API::port", port)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	return
//}
