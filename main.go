package main

import (
	"os"
)

func main() {
	command := os.Args[1]
	switch command {
	case "workshop":
		NotifyWorkshopInfo(GetWorkshopInfo(os.Args[2]))
	case "user":
		NotifyUserInfo(GetUserInfo(os.Args[2]))
	case "server":
		StartServer(os.Args[2])
	}
}
