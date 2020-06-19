package main

import (
	"fmt"
	"os"

	"gopkg.in/gookit/color.v1"
)

func main() {
	command := os.Args[1]
	switch command {
	case "workshop":
		NotifyWorkshopInfo(GetWorkshopInfo(os.Args[2]))
	case "user":
		NotifyUserInfo(GetUserInfo(os.Args[2]))
	case "server":
		fmt.Print("Server hosts on ")
		color.Print("<fg=blue;op=underscore;>http://</>")
		color.Print("<fg=cyan;op=underscore;>127.0.0.1:</>")
		color.Printf("<fg=green;op=underscore;>%s\n</>", os.Args[2])
		color.Print("<op=reverse>[Cotrol]</>+<op=reverse>[C]</> to break")
		StartServer(os.Args[2])
	}
}
