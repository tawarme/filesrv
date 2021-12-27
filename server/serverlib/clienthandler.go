package serverlib


import (
        "fmt"
        "net"
)


func ClientHandler(client net.Conn) {
        for true {
                buf := make([]byte, 1024)
                _, err := client.Read(buf)

                if err != nil { 
                        fmt.Println(err)
                        return
                }

                fmt.Println("Client", client.RemoteAddr(), ", received data: ", string(buf))

                command := string(buf[:3])
                fmt.Println("Client", client.RemoteAddr(), ", received command: ", command)

                switch command {
                case "SUB":
                        fmt.Println("Client", client.RemoteAddr(), ", subbing to: ")
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
