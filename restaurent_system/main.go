package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/devika-singh/GO/GO/restaurent_system/pkg/login"
	"github.com/devika-singh/GO/GO/restaurent_system/pkg/manager"
)

func main() {
	loginInterface := login.Login{}
	loginInterface.GetUserInfo()
	for {
		user := loginInterface.LoginPage()
		switch strings.ToLower(user) {
		case "server":
			fmt.Println("Server")
		case "manager":
			manager.ManagerInterface()
		case "exit":
			return
		default:
			fmt.Println("Enter valid user")
			time.Sleep(2 * time.Second)
		}
	}
}
