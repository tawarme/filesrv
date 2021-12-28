package main


import (
	"fmt"
	"example.com/filesrv/client/clientlib"
)


func main() {
	server_add := "localhost"

	fmt.Println("Client started. Server is:", server_add)
	clientlib.Client(server_add)
}
