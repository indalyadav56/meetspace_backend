package main

import (
	"fmt"
	"net"
)

func main() {
	// Define SMTP server configuration
	smtpHost := "0.0.0.0" // Change this to your server's IP address or domain name
	smtpPort := "25"      // SMTP default port

	// Start listening on the SMTP port
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", smtpHost, smtpPort))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	fmt.Println("SMTP server is listening on", smtpHost+":"+smtpPort)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle client connection in a separate goroutine
		go handleClient(conn)
	}
}

// Handle client connection
func handleClient(conn net.Conn) {
	defer conn.Close()

	// Send SMTP server greeting message
	conn.Write([]byte("220 SMTP Server Ready\r\n"))

	// Read client command
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Respond to client command
	command := string(buffer)
	if command[0:4] == "HELO" {
		conn.Write([]byte("250 Hello\r\n"))
	} else if command[0:4] == "QUIT" {
		conn.Write([]byte("221 Goodbye\r\n"))
	} else {
		conn.Write([]byte("500 Command not recognized\r\n"))
	}
}
