package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/senseyeio/roger"
)

func prompt(rClient roger.RClient) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("quit", text) == 0 {
			break
		}
		send_command(text, rClient)
	}
}

func send_command(command string, rClient roger.RClient) {
	value, err := rClient.Eval(command)
	if err != nil {
		fmt.Println("Command failed: " + err.Error())
	} else {
		fmt.Println(value) // 3.141592653589793
	}
}

// https://www.rforge.net/Rserve/doc.html#cmdl
func main() {
	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}

	prompt(rClient)
}
