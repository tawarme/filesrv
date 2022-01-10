package main


import (
	"fmt"
	"example.com/filesrv/client/clientlib"
	"flag"
	"os"
)


func main() {
	server_add := "localhost"

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	sendChannel := sendCmd.Int("channel", 1, "Channel to send file.")

	receiveCmd := flag.NewFlagSet("send", flag.ExitOnError)
	receiveChannel := receiveCmd.Int("channel", 1, "Channel to send file.")

	if len(os.Args) < 2 {
		fmt.Println("Missing mode")
		fmt.Println("Valid modes: send receive")
		return
	}
	mode := os.Args[1]

	var channel *int
	file_path := ""

	switch mode {
	case "send":
		if len(os.Args) < 3 {
			fmt.Println("Missing file path")
			return
		}

		file_path = os.Args[2]

		sendCmd.Parse(os.Args[3:])
		channel = sendChannel

		break

	case "receive":
		receiveCmd.Parse(os.Args[2:])
		channel = receiveChannel
		break

	default:
		fmt.Println("Invalid mode:", mode)
		fmt.Println("Valid modes: send receive")
		return
	}

	fmt.Println("Client started. Server is:", server_add)
	clientlib.Client(server_add, mode, *channel, file_path)
}
