package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Request message to the server
type Request = string

const help = `
Commands:

	set KEY VAL         Set the key-value pair
	get KEY             Get value by key
	del KEY             Delete value by key
	all                 Dump database
	clr                 Clear database
	siz                 Get number of pairs

	:q                  Quit the client
	:h                  Show help message
`

var (
	errInvalidCommand = errors.New("Invalid command")
)

// ParseCmd parses the command and gives out request
func ParseCmd(txt string) (Request, error) {
	strs := strings.Fields(txt)
	cmd := strings.ToLower(strs[0])
	args := strs[1:]

	switch len(args) {
	case 2:
		if cmd == "set" {
			return fmt.Sprintf("+%v\r\n%v", args[0], args[1]), nil
		}
		return "", errInvalidCommand

	case 1:
		var sym string
		switch cmd {
		case "get":
			sym = "="
		case "del":
			sym = "-"
		default:
			return "", errInvalidCommand
		}
		return sym + args[0], nil

	case 0:
		switch cmd {
		case "all":
			return "*", nil
		case "clr":
			return "!", nil
		case "siz":
			return "?", nil
		}

	}

	return "", errInvalidCommand
}

// ShowHelp displays help message
func ShowHelp() {
	fmt.Println(help)
}

func main() {
	port := flag.Int("p", 8080, "port to dial")
	addr := fmt.Sprintf(":%d", *port)

	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("dictd> ")
		rawtxt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		txt := strings.TrimSpace(rawtxt)

		if txt == ":q" {
			c.Close()
			break
		}
		if txt == ":h" {
			ShowHelp()
			continue
		}

		cmd, err := ParseCmd(txt)

		if err != nil {
			log.Println(err)
			ShowHelp()
			continue
		}

		fmt.Fprintf(c, cmd+"\r\n")

		res, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)
	}
}
