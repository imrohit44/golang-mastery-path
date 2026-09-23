package main

import "fmt"

// --- Simulating a different package named 'models' ---
type UserAccount struct {
	Username string // Exported: Accessible anywhere
	password string // Unexported: Accessible ONLY inside this package
}

func NewUser(user, pass string) UserAccount {
	return UserAccount{
		Username: user,
		password: pass,
	}
}

// CheckPassword has access to the unexported password field
func (u UserAccount) CheckPassword(attempt string) bool {
	return u.password == attempt
}
// -----------------------------------------------------

func main() {
	user := NewUser("rohit_admin", "supersecret123")
	
	fmt.Println("Username:", user.Username)
	
	// user.password is NOT accessible here if this was actually a separate package.
	// We use the exported method instead:
	isValid := user.CheckPassword("supersecret123")
	fmt.Println("Password valid?", isValid)
}