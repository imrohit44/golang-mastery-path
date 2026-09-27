package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	// Create or open a file
	file, err := os.OpenFile("lockfile.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer os.Remove("lockfile.txt")

	fmt.Println("Attempting to acquire exclusive OS lock on file...")

	// syscall.LOCK_EX = Exclusive Lock, syscall.LOCK_NB = Non-Blocking
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		fmt.Println("Failed to acquire lock. Another process holds it:", err)
		return
	}

	fmt.Println("Lock acquired successfully! Simulating work...")
	time.Sleep(2 * time.Second)

	// Release the lock
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	if err != nil {
		fmt.Println("Failed to release lock:", err)
	} else {
		fmt.Println("Lock released.")
	}
}