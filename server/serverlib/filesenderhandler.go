package serverlib


import (
	"fmt"
	"net"
	"os"
	"bytes"
)


func FileSenderHandler(clients_list []net.Conn, client net.Conn, headers []byte, file_name string) {
    fmt.Println("Sending to client", client.RemoteAddr())
	
	f, err := os.Open(file_name)

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


	data := make([]byte, 1024-(len(headers)+1))
	_, err = f.Read(data)

	if err != nil {
		fmt.Println(err)
		return 
	}

	buf := bytes.Join([][]byte{ headers,
								data },
					  []byte(" "))

	client.Write(buf)

	transmitted_so_far := 1024-(len(headers)+1)
	for transmitted_so_far < int(file_length) {
		data := make([]byte, 1024)
		count, err := f.Read(data)

		if err != nil {
			fmt.Println(err)
			return 
		}

		client.Write(data)

		transmitted_so_far += count
	}
	client.Close()

	fmt.Println("Finished sending file.")

}
