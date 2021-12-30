package clientlib

import (
	"fmt"
	"os"
	"net"
	"encoding/binary"
	"bytes"
)

func ReceiverHandler(server net.Conn, channel int) {
	// Subscribe to channel

	channel_enc := make([]byte, 4)
	binary.BigEndian.PutUint32(channel_enc, uint32(channel))

	headers := bytes.Join([][]byte{ []byte("SUB"), 
									channel_enc},
					  	  []byte(" "))

	server.Write(headers)

	buf := make([]byte, 1024)

	_, err := server.Read(buf)

	if err != nil { 
        fmt.Println(err)
        return
    }

    fmt.Println("Received data:", buf)

    // 0         1         2
    // 012345678901234567890123456789
    // PUT X AAAAAA YYYY BBBBBBBBBBBBBB
    // name_length is 1 byte unsigned int (X)
    // name is as many bytes as name_length (AAAAAA)
    // content_length is 4 bytes (YYYY)
    // content is as many bytes as content_length (BBBBBBBBBBBBBB)

    name_length := int(buf[4])

    file_name := string(buf[6:6+name_length])

    content_length_offset := 6+name_length+1
    content_length := int(binary.BigEndian.Uint32(buf[content_length_offset:content_length_offset + 4]))

    content_offset := content_length_offset+4 +1
    
    fmt.Println("Putting file: ", file_name, "size: ", content_length)

    headers_length := 3+1+ 1+1+ name_length+1+ 4+1

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

            count, err := server.Read(data)

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
            stats, _ := f.Stat()

            received_so_far += count
    }
    f.Sync()
    server.Close()
    fmt.Println("Finished receiving file.")

}
