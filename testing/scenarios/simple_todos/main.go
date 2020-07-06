package main

import "fmt"

// TODO: This is a malformed todo
// TODO: This is a malformed todo 2

// This is not a todo comment and shouldn't be matched

func main() { // TODO: This is a todo comment at the end of a line

	fmt.Println("This is a TODO string, which shouldn't be matched")
}

// TODO comment without colons

// This is a TODO comment at the middle of it
