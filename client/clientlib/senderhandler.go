package clientlib

import (
	"fmt"
	//"os"
	"net"
)
func SenderHandler(server net.Conn, channel int, file_path string) {
	fmt.Println("Sending to channel", channel, "file", file_path)
}
