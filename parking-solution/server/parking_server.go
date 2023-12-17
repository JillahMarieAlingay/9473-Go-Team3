package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	RunServer()
}

func RunServer() {
	// Create a server variable 'listener'
	listener, err := net.Listen("tcp", "localhost:8080")
	// Exception handling
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	// Construct server status message
	serverStatus := fmt.Sprint("Server is running on ", listener.Addr())

	// Print server status message
	fmt.Println(serverStatus)

	for {
		// Accept connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue // Continue listening for other clients
		}

		// Goroutine = Java (concurrency)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Schedule the connection to close
	defer conn.Close()

	for {
		// 'Byte slice' to store data read
		buffer := make([]byte, 1024)
		// Read data from the connection
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			break // Break for loop if there is an error
		}

		// Convert the received data to a string
		clientMessage := string(buffer)

		clientMessagePrint := fmt.Sprint("Received message from client:", clientMessage)

		// Print client's message to the console
		fmt.Println(clientMessagePrint)

		// Read data from the text file (BufferedReader)
		/* BufferedReader bufferedReader = new BufferedReader(new FileReader(fileName));

		   StringBuilder parkingData = new StringBuilder();
		   String line;
		   while ((line = bufferedReader.readLine()) != null) {
		       parkingData.append(line).append("\n");
		   } */
		parkingData, err := os.ReadFile("parking_data.txt")
		if err != nil {
			fmt.Println("Error reading data from file:", err)
			break
		}

		// Send the parking data as a response to the client
		_, err = conn.Write(parkingData)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}

		// Close the connection after sending the response
		conn.Close()
		break
	}
}
