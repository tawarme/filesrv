package serverlib


import (
        "fmt"
        "net"
)


func Server() {
        clients_subscriptions := make(map[string]uint32)

	listener, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
            fmt.Println(err)
            return
        }

	for true {
		
                client, err := listener.Accept()

                if err != nil {
                        fmt.Println(err)
                        return
                }

                fmt.Println("Connected to client: ", client.RemoteAddr())

                go ClientHandler(clients_subscriptions, client)

	}
}
