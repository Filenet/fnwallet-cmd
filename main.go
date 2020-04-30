package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var i = make(chan int)

func main() {
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	Help(nil)
	go func() {
		for {
			b, _, _ := r.ReadLine()
			line := string(b)
			tokens := strings.Split(line, " ")
			if _, ok := handlers[tokens[0]]; ok {
				handlers[tokens[0]](nil)
			} else {
				Help(nil)
			}
		}
	}()
	<-i
}
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"help":    Help,
		"h":       Help,
		"version": Vsersion,
		"v":       Vsersion,
		"quit":    Quit,
		"exit":    Quit,
	}
}

func Help(args []string) int {
	fmt.Println(`These are common tool commands used in various situations:
       version  this is tool version
       exit     quit tool
       quit     quit tool
       `)
	return 0
}

func Quit(args []string) int {
	i <- 1
	return 0
}

func Vsersion(args []string)int{
	fmt.Println(`version 1.0.0`)
	return 0
}