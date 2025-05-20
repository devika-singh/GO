package login

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/devika-singh/GO/GO/restaurent_system/pkg/io_module"
	"github.com/devika-singh/GO/GO/restaurent_system/pkg/user"
)

type Login struct {
	registeredUserList []user.User
}

func (l *Login) GetUserInfo() {
	f, err := os.Open("user/users.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	l.registeredUserList = l.registeredUserList[:0] // reset slice

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		/* discarded header*/
	}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // skip blank lines
		}

		// Split by any run of spaces (simplest for the padded format).
		fields := strings.Fields(line)
		if len(fields) < 3 {
			fmt.Printf("skipping malformed line: %q\n", line)
			continue
		}

		name := fields[0]

		id, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("bad id on line %q: %v\n", line, err)
			continue
		}

		role := fields[2]

		l.registeredUserList = append(
			l.registeredUserList,
			user.User{Name: name, ID: id, Role: role},
		)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(l.registeredUserList)
}

func (l *Login) LoginPage() string {
	io_module.ClearScreen()
	r := bufio.NewReader(os.Stdin)

	role, _ := io_module.GetInput("Enter your role or exit: ", r)

	return role
}
