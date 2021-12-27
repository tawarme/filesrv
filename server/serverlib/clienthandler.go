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

                fmt.Println("Client", client.RemoteAddr(), ", received: ", string(buf))
        }
}
