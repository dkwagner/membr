package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {

	commandFlag := flag.Bool("command", false, "Create new membr entry as a command")
	textFlag := flag.Bool("text", false, "Create new membr entry as text")
	linkFlag := flag.Bool("link", false, "Create new membr entry as a link")

	flag.Parse() // Parses flags, if not done, no flags are parsed

	// We only want to allow one command to be set, otherwise error out
	command, err := flagSet([]*bool{commandFlag, textFlag, linkFlag})

	if err != nil {
		fmt.Println(err)
		return
	}

	// var input string

	switch command {
	case 0:
		fmt.Println("Membr type: command")
	case 1:
		fmt.Println("Membr type: text")
	case 2:
		fmt.Println("Membr type: link")
	}
}

func flagSet(flags []*bool) (int, error) {

	flag := -1
	numFlagsSet := 0

	for i := 0; i < len(flags); i++ {
		if *flags[i] {
			flag = i
			numFlagsSet++
		}
	}

	// If more than one flag set, error out
	if numFlagsSet != 1 {
		return -1, errors.New("ERROR: You can specify only one Membr creation type")
	}

	return flag, nil
}
