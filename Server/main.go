package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 从命令行接收IP地址
	if len(os.Args) < 2 {
		fmt.Println("Usage: server.exe <IP_ADDRESS>")
		return
	}
	ipAddress := os.Args[1]
	port := "9999"

	conn, err := net.Dial("tcp", ipAddress+":"+port)
	if err != nil {
		fmt.Println("Error connecting to client:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to client. Enter commands to execute.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		if command == "exit\n" {
			fmt.Println("Closing connection...")
			break
		}

		_, err := conn.Write([]byte(command))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		fmt.Println("Command output:\n", string(buffer[:n]))
	}

	fmt.Println("Connection closed.")
}
