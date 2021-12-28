package serverlib


import (
        "fmt"
        "net"
        "github.com/nu7hatch/gouuid"
        "encoding/binary"
)


func ClientHandler(clients_subscriptions map[string]uint32, client net.Conn,) {
        for true {
                u, err := uuid.NewV4()
                
                if err != nil { 
                        fmt.Println(err)
                        return
                }             

                id := u.String()
                clients_subscriptions[id] = 0

                buf := make([]byte, 1024)
                _, err = client.Read(buf)

                if err != nil { 
                        fmt.Println(err)
                        return
                }

                fmt.Println("Client", client.RemoteAddr(), ", received data: ", string(buf))

                command := string(buf[:3])
                fmt.Println("Client", client.RemoteAddr(), ", received command: ", command)

                switch command {
                case "SUB":
                        // SUB XXXX
                        // channel is 4 bytes unsigned int(XXXX)"

                        channel := binary.BigEndian.Uint32(buf[4:8])
                        clients_subscriptions[id] = channel

                        fmt.Println("Client", client.RemoteAddr(), ", subbing to: ", channel)
                        break
                case "PUT":
                        fmt.Println("Client", client.RemoteAddr(), ", putting file: ")
                        break
                case "GET":
                        fmt.Println("Client", client.RemoteAddr(), ", getting file: ")
                        break
                case "UNK":
                        fmt.Println("Client", client.RemoteAddr(), ", command not understood.")
                        break
                default:
                        fmt.Println("Client", client.RemoteAddr(), ", unrecognized command: ", command)
                }
        }
}
