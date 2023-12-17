package main

import (
	"fmt"
	"net"
)

func main() {
	RunClient()
}

func RunClient() {
	// Establish a TCP connection to localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	// Schedule the connection to close when the method exits
	defer conn.Close()

	// Send byte array message to server
	/*	OutputStream outputStream = conn.getOutputStream(); // Assuming 'conn' is a Socket or OutputStream
		String message = "Parking status request. (801)";
		byte[] messageBytes = message.getBytes("UTF-8"); // Convert the string to bytes with the appropriate encoding
		outputStream.write(messageBytes);
		outputStream.flush(); // Ensure that the data is sent immediately */
	message := "Parking status request. (801)"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	// Store data from a 'byte slice'
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		return
	}

	fmt.Println(string(response[:n]))
}
