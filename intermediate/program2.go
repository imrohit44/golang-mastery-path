package main

import (
	"fmt"
	"time"
)

type BusinessError struct {
	Code    int
	Message string
	When    time.Time
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("[%s] Error %d: %s", e.When.Format(time.RFC3339), e.Code, e.Message)
}

func processOrder(amount float64) error {
	if amount <= 0 {
		return &BusinessError{
			Code:    400,
			Message: "Order amount must be greater than zero",
			When:    time.Now(),
		}
	}
	return nil
}

func main() {
	if err := processOrder(-15.5); err != nil {
		fmt.Println("Failed:", err)
	}
}