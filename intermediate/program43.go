package main

import (
	"errors"
	"fmt"
	"os"
)

// A custom error type
type DatabaseError struct {
	Query string
	Msg   string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("DB Error [%s]: %s", e.Query, e.Msg)
}

func executeQuery() error {
	err := &DatabaseError{Query: "SELECT *", Msg: "connection timeout"}
	// Wrap the error with context
	return fmt.Errorf("failed to execute query: %w", err)
}

func main() {
	err := executeQuery()

	// 1. Check for specific error instances (errors.Is)
	if errors.Is(err, os.ErrPermission) {
		fmt.Println("Permission denied.")
	}

	// 2. Extract a specific error type (errors.As)
	var dbErr *DatabaseError
	if errors.As(err, &dbErr) {
		fmt.Printf("Recovered original DB error. Query was: %s\n", dbErr.Query)
	}

	fmt.Println("Full error trace:", err)
}