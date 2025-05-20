package user

import (
	"testing"
)

// ---------- UserRow & CreateUser ----------

func TestUserRow(t *testing.T) {
	var u User
	u.CreateUser("Alice", 7, "admin")

	want := "Alice                     7  admin     \n" // 25-2-10 format width
	got := u.UserRow()

	if got != want {
		t.Errorf("UserRow() mismatch:\nwant %q\ngot  %q", want, got)
	}
}

func TestSaveUser(t *testing.T) {
	var u User
	u.CreateUser("Bob", 42, "server")
	u.SaveUser()

	// 5. call SaveUser again and ensure it appends
	u.CreateUser("Carol", 9, "manager")
	u.SaveUser()
}
