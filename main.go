package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	done := make(chan bool)

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Enter a port number.")
		return
	}

	if len(flag.Args()) == 1 {
		fmt.Print("Listening on port: ")
	} else {
		fmt.Print("Listening on ports: ")
	}

	for _, port := range flag.Args() {
		fmt.Printf("%v ", port)
		go listen(port, done)
	}

	<-done
}

func listen(port string, end chan bool) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		end <- true
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			end <- true
			return
		}
		conn.Write([]byte("Port is open. Goodbye!\n"))
		conn.Close()
	}
}
