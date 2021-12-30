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

	_, file_name := filepath.Split(file_path)


	channel_enc := make([]byte, 4)
	binary.BigEndian.PutUint32(channel_enc, uint32(channel))


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
	_, err = f.Read(data)

	if err != nil {
		fmt.Println(err)
		return 
	}

	buf := bytes.Join([][]byte{ headers,
								data },
					  []byte(" "))

	server.Write(buf)

	transmitted_so_far := 1024-(len(headers)+1)
	for transmitted_so_far < int(file_length) {
		data := make([]byte, 1024)
		count, err := f.Read(data)

		if err != nil {
			fmt.Println(err)
			return 
		}

		server.Write(data)

		transmitted_so_far += count
	}
	server.Close()

	fmt.Println("Finished sending file.")
}
