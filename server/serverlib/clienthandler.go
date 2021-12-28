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
                        // PUT X AAAAAA YYYY BBBBBBBBBBBBBB
                        // name_length is 1 byte unsigned int (X)
                        // name is as many bytes as name_length (AAAAAA)
                        // content_length is 4 bytes (YYYY)
                        // content is as many bytes as content_length (BBBBBBBBBBBBBB)

                        name_length := buf[4]
                        file_name := string(buf[5:5+name_length])

                        content_length_offset = 5+name_length+1
                        content_length := buf[content_length_offset:content_length_offset + 4]

                        content = buf[content_length_offset + 4 +1:]


                        fmt.Println("Client", client.RemoteAddr(), ", putting file: ", file_name, "size: ", content_length)


                        f, err := os.Create(file_name)
                        
                        if err != nil {
                                fmt.Println(err)
                                return
                        }

                        defer f.Close()

                        _, err := f.Write(content)

                        if err != nil {
                                fmt.Println(err)
                                return
                        }

                        f.sync()

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
