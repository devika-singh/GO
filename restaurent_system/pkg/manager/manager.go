package manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/devika-singh/GO/GO/restaurent_system/pkg/io_module"
)

func ManagerInterface() {
	fmt.Println("Manager Interface")

	r := bufio.NewReader(os.Stdin)

	for {
		opt, _ := io_module.GetInput("Select Option [m : \"Edit Menu\"/ u : \"Edit Users\"] :/ e : \"exit\"", r)

		switch strings.ToLower(opt) {
		case "m":
			UpdateMenu(r)
		case "u":
			UpdateUsers(r)
		case "e":
			return
		default:
			fmt.Println("Enter valid option")
		}
	}
}

func UpdateMenu(r *bufio.Reader) {
	for {
		opt, _ := io_module.GetInput("Select Option [a : \"Add item\"/ r : \"Remove item\"/ t : \"Toggle item status\"] :", r)
		switch opt {
		case "a":
			name, _ := io_module.GetInput("Enter item name: ", r)
			price, _ := io_module.GetInput("Enter item price: ", r)
			// menu.AddItem(name, price)
			fmt.Println("Add item", name, price)
		}
	}
}

func UpdateUsers(r *bufio.Reader) {
	for {
		opt, _ := io_module.GetInput("Select Option [a : \"Add User\"/ r : \"Remove User\"] :", r)
		switch opt {
		case "a":
			id, _ := io_module.GetInput("Enter user id: ", r)
			// menu.AddItem(name, price)
			fmt.Println("Add item", id)
		}
	}

}
