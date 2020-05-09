package cmd

import (
	"bufio"
	"fmt"
	"fnv3/test/filenetipfs"
	"github.com/astaxie/beego/config"
	"os"
	"strings"
)

var Chanl = make(chan int)

func init() {
	Filenet["Help"] = FilenetHelp
	Filenet["Add"] = FilenetAdd
	Filenet["get"] = FilenetGet
	FilenetConfig, err := config.NewConfig("ini", "filnetipfs.config")
	if err != nil {
		panic(err)
	}
	InitIpfsPort(FilenetConfig)
	port:="8081"
	fmt.Println(FilenetConfig.Set("API::port",port))

}


func InitIpfsPort(con config.Configer) {
	port := con.String("API::port")
	filenetipfs.SetIpfsAddPath(port)
	filenetipfs.SetIpfsLsPath(port)
	filenetipfs.SetIpfsBlockRawPath(port)
}


func Command() {
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
		case "filenet":
			FilenetServer(comm[1:])
		case "bye":
			FilenetExit(nil)
		case "exit":
			FilenetExit(nil)
		default:
			res:=fmt.Sprintf("%s\t%s",COMMANDNOTFOUND,comm[0])
			os.Stderr.Write([]byte(res))
		}
	}
}

func FilenetExit(args []string){
	<-Chanl
	return
}
