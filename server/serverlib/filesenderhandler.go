package serverlib


import (
	"fmt"
	"net"
)

func FileSenderHandler(clients_list []net.Conn, client net.Conn) {
    fmt.Println("Sending to client", client.RemoteAddr())
	
}
