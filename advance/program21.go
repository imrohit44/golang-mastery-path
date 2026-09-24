package main

/*
#include <stdio.h>
#include <math.h>

void sayHelloFromC() {
    printf("Hello from the C runtime!\n");
}

double calculateSquareRoot(double x) {
    return sqrt(x);
}
*/
import "C"
import "fmt"

func main() {
	// Call a void C function
	C.sayHelloFromC()

	// Pass Go variables to C and get the result
	input := 144.0
	result := C.calculateSquareRoot(C.double(input))

	fmt.Printf("Go sees the square root of %.2f as: %.2f\n", input, result)
}