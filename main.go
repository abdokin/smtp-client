package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run main.go <host> <port> <username> <password>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]
	username := os.Args[3]
	password := os.Args[4]

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal("Error connecting to the server:", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Read the response.
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	fmt.Println(response)
	// Send the AUTH PLAIN command with base64 encoded credentials.
	authString := "\x00" + username + "\x00" + password
	authStringBase64 := base64.StdEncoding.EncodeToString([]byte(authString))

	fmt.Fprintf(writer, "AUTH PLAIN %s\r\n", authStringBase64)
	writer.Flush()
	println("Trying to AUTH");

	// Read the response.
	response, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	fmt.Println(response)
}
