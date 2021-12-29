package clientlib

import (
	"fmt"
	"os"
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

	f, err := os.Open(file_path)

	if err != nil {
		fmt.Println(err)
		return 
	}

	file_stats, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	file_length := file_stats.Size()

	content_length := make([]byte, 4)
	binary.BigEndian.PutUint32(content_length, uint32(file_length))

	headers := bytes.Join([][]byte{ []byte("PUT"), 
								channel_enc,
								name_length, 
								[]byte(file_name), 
								content_length },
					  	  []byte(" "))

	data := make([]byte, 1024-(len(headers)+1))
	count, err := f.Read(data)

	if err != nil {
		fmt.Println(err)
		return 
	}

	buf := bytes.Join([][]byte{ headers,
								data },
					  []byte(" "))

	server.Write(buf)
}
