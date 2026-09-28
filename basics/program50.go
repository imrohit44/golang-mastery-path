package main

import "fmt"

func main() {
	i := 0

StartLoop: // Label
	if i >= 5 {
		goto End // Jump to End
	}
	
	fmt.Println("Counter:", i)
	i++
	goto StartLoop // Jump back to StartLoop

End:
	fmt.Println("Loop finished using goto.")
}