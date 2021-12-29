package clientlib

import (
	"fmt"
	//"os"
	"net"
	"path/filepath"
	"encoding/binary"
	"bytes"
)


func SenderHandler(server net.Conn, channel int, file_path string) {
	fmt.Println("Sending to channel", channel, "file", file_path)

	dir, file_name := filepath.Split(file_path)
	fmt.Println(dir, file_name)

	channel_enc := make([]byte, 4)
	binary.BigEndian.PutUint32(channel_enc, uint32(channel))

	fmt.Println(channel_enc)

	name_length := make([]byte, 1)
	name_length[0] = uint8(len(file_name))

	// dummy vals
	content_length := make([]byte, 4)
	binary.BigEndian.PutUint32(content_length, uint32(5))

	content := "AEIOU"

	buf := bytes.Join([][]byte{ []byte("PUT"), 
								channel_enc,
								name_length, 
								[]byte(file_name), 
								content_length,
								[]byte(content)
							   },
					  []byte(" "))

	server.Write(buf)
}
