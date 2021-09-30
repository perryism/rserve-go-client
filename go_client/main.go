package main

import (
	"bufio"
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

func NewPrompt(host string, port int) RServeConnect {

	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		panic("Failed to connect")
	}
	return RServeConnect{rClient: rClient}
}

// https://www.rforge.net/Rserve/doc.html#cmdl
func main() {
	conn := NewPrompt("127.0.0.1", 6311)
	conn.prompt()
}
