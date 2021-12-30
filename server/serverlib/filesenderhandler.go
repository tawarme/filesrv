package serverlib


import (
	"fmt"
	"net"
)

func FileSenderHandler(clients_list []net.Conn, client net.Conn, headers []byte, file_name string) {
    fmt.Println("Sending to client", client.RemoteAddr())
	
	client.Write(headers)
}
