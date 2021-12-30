package serverlib


import (
        "fmt"
        "net"
        "encoding/binary"
        "os"
)


func ClientHandler(clients_subscriptions map[uint32][]net.Conn, client net.Conn,) {
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
                        // SUB XXXX
                        // channel is 4 bytes unsigned int(XXXX)"

                        channel := binary.BigEndian.Uint32(buf[4:8])
                        clients_subscriptions[channel] = append(clients_subscriptions[channel], client)

                        fmt.Println("Client", client.RemoteAddr(), ", subbing to: ", channel)
                        // We will only send, not receive info
                        return
                        break
                case "PUT":
                        // 0         1         2
                        // 012345678901234567890123456789
                        // PUT ZZZZ X AAAAAA YYYY BBBBBBBBBBBBBB
                        // channel is 4 byte unsigned int(ZZZZZ)
                        // name_length is 1 byte unsigned int (X)
                        // name is as many bytes as name_length (AAAAAA)
                        // content_length is 4 bytes (YYYY)
                        // content is as many bytes as content_length (BBBBBBBBBBBBBB)

                        channel := binary.BigEndian.Uint32(buf[4:8])

                        name_length := int(buf[9])

                        file_name := string(buf[11:11+name_length])

                        content_length_offset := 11+name_length+1
                        content_length := int(binary.BigEndian.Uint32(buf[content_length_offset:content_length_offset + 4]))

                        content_offset := content_length_offset+4 +1
                        
                        fmt.Println("Client", client.RemoteAddr(), "to channel", channel, "putting file: ", file_name, "size: ", content_length)

                        headers_length := 3+1+ 4+1+ 1+1+ name_length+1+ 4+1

                        var content_segment int
                        if content_length >= 1024 - headers_length {
                                content_segment = 1024 - headers_length
                        } else {
                                content_segment = content_length
                        }


                        content := buf[content_offset:content_offset+content_segment]

                        f, err := os.Create(file_name)
                        
                        if err != nil {
                                fmt.Println(err)
                                return
                        }

                        defer f.Close()

                        _, err = f.Write(content)

                        if err != nil {
                                fmt.Println(err)
                                return
                        }

                        f.Sync()

                        received_so_far := content_segment

                        for received_so_far < content_length {
                                data := make([]byte, 1024)

                                count, err := client.Read(data)

                                if err != nil { 
                                        fmt.Println(err)
                                        return
                                }

                                if content_length < (received_so_far + count) {
                                        _, err = f.Write(data[:content_length-received_so_far])
                                } else {
                                        _, err = f.Write(data)
                                }

                                if err != nil {
                                        fmt.Println(err)
                                        return
                                }
                                
                                f.Sync() 

                                received_so_far += count
                        }
                        client.Close()

                        for _, client := range clients_subscriptions[channel] {
                                go FileSenderHandler(clients_subscriptions[channel], client)
                        }

                        return
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
