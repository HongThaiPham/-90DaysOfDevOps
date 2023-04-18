package main

import "fmt"

func main() {
	var challenge = "#90DaysOfDevOps"
	const daystotal = 90
	var dayscomplete = 11

	fmt.Println("Welcome to", challenge, "")
	fmt.Println("This is a", daystotal, "challenge and you have completed", dayscomplete, "days")
	fmt.Printf("This is a %v challenge and you have completed %v days\n", daystotal, dayscomplete)
	fmt.Println("Great work")
}
