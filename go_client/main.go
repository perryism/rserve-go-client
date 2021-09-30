package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/senseyeio/roger"
)

type RServeConnect struct {
	rClient roger.RClient
}

func (conn *RServeConnect) sendCommand(command string) {
	value, err := conn.rClient.Eval(command)
	if err != nil {
		fmt.Println("Command failed: " + err.Error())
	} else {
		fmt.Println(value)
	}
}

func (conn *RServeConnect) prompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("quit", text) == 0 {
			break
		}
		conn.sendCommand(text)
	}
}

func NewPrompt(host string, port int64) RServeConnect {

	rClient, err := roger.NewRClient(host, port)
	if err != nil {
		panic("Failed to connect")
	}
	return RServeConnect{rClient: rClient}
}

// https://www.rforge.net/Rserve/doc.html#cmdl
func main() {
	port := flag.Int64("port", 6311, "port")
	host := flag.String("host", "127.0.0.1", "host")
	cmmd := flag.String("c", "", "Command")
	flag.Parse()

	conn := NewPrompt(*host, *port)

	if len(*cmmd) == 0 {
		conn.prompt()
	} else {
		conn.sendCommand(*cmmd)
	}
}
