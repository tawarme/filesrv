package main


import (
	"fmt"
	"example.com/filesrv/client/clientlib"
	"flag"
)


func main() {
	server_add := "localhost"


	channel := flag.Int("channel", 1, "Channel to send file.")

	flag.Parse()

	mode := flag.Arg(0)

	file_path := ""
	switch mode {
	case "send":
		file_path = flag.Arg(1)
		if file_path == "" {
			fmt.Println("Missing file path")
			return
		}
		break
	case "receive":
		break
	default:
		fmt.Println("Invalid mode:", mode)
		fmt.Println("Valid modes: send receive")
		return
	}

	fmt.Println("Client started. Server is:", server_add)
	clientlib.Client(server_add, mode, *channel, file_path)
}
