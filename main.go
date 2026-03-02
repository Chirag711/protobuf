package main

import (
	"fmt"
	"mainprotobuf/userpb"
)

func main() {
	// This is a placeholder main function. The actual logic for using the generated protobuf code would go here.
	user := userpb.User{
		Id:   1,
		Name: "Chirag",
	}

	fmt.Println(user.Id)
	fmt.Println(user.Name)
}
