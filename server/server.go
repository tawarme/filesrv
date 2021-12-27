package main


import (
	"fmt"
	"example.com/filesrv/server/serverlib"
)


func main() {
	fmt.Println("Server started, listeing on port 5000.")
	serverlib.Server()
}
