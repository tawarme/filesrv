package clientlib

import (
	"fmt"
	//"os"
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

}
