package clientlib


import (
        "fmt"
        "net"
)


func Client(server_add string, mode string, channel int, file_path string) {
	server, err := net.Dial("tcp", server_add+":5000")

	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println("Connected to server")

	switch mode {
	case "send":
		SenderHandler(server, channel, file_path)
		break
	case "receive":
		ReceiverHandler(server, channel)
	default:
		/* code */
		return
	}
}
