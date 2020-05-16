package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	API     APIInfo     `json:"api,omitempty"`
	Block   BlockConfig `json:"block,omitempty"`
	Library LibraryInfo `json:"library,omitempty"`
}

type LibraryInfo struct {
	Path string `json:"path,omitempty"`
}

type BlockConfig struct {
}

type APIInfo struct {
	Ipfs_port string `json:"ipfsport,omitempty"`
	Http_port string `json:"http_port,omitempty"`
}

func init() {
	file, err := os.OpenFile("./config.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if fileInfo.Size() == 0 {
		config := genesisConfig()
		configByte, err := json.MarshalIndent(config," "," ")
		if err != nil {
			panic(err)
		}
		_, err = file.Write(configByte)
		if err != nil {
			panic(err)
		}
	}
}

func genesisConfig() *Config {
	config := new(Config)
	config.API.Ipfs_port = "5001"
	config.API.Http_port = "8081"
	config.Library.Path = "/leveldb/root/"
	return config
}

func (c *Config) Save() error {
	file, err := os.OpenFile("./config.json", os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := json.MarshalIndent(c," "," ")
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func NewConfig() (*Config, error) {
	conf, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}
	var config = new(Config)
	err = json.Unmarshal(conf, config)
	if err != nil {
		return nil, err
	}
	return config, err
}

func (c *Config) ConfigMap() (*map[string]interface{}, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	var configMap = make(map[string]interface{})
	err = json.Unmarshal(b, &configMap)
	if err != nil {
		return nil, err
	}
	return &configMap, nil
}

func (c *Config) String() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func Set(args string)(*Config, error){
	return set(args)
}

func set(args string) (*Config, error) {
	commad := strings.Split(args, "=")
	config, err := parseCommand(commad[0], commad[1])
	if err != nil {
		return nil, err
	}
	err = config.Save()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func parseCommand(args string, value string) (*Config, error) {
	strs := strings.Split(args, ".")
	var format string
	for _, s := range strs {
		format = fmt.Sprintf("%s\"%s\":", format, s)
	}
	str := strings.ReplaceAll(format, ":", ":{")
	n := strings.Count(str, "{")
	var j string
	for i := 0; i < n; i++ {
		j = fmt.Sprintf("%s%s", j, "}")
	}
	l := fmt.Sprintf("{%s%s}", str, j)
	var config = new(Config)
	c := strings.ReplaceAll(l, "{}", "\""+value+"\"")
	err := json.Unmarshal([]byte(c), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
