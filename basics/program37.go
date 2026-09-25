package main

import "fmt"

func main() {
	users := map[int]string{
		101: "Alice",
		102: "Bob",
		103: "Charlie",
	}

	fmt.Println("Before deletion:", users)

	// Delete user 102
	delete(users, 102)

	fmt.Println("After deletion:", users)
	
	// Deleting a non-existent key is safe and does nothing
	delete(users, 999) 
}