package main

import (
	"errors"
	"fmt"
	"os"
	"whatsappminer/lib"
)

func main() {
	path := os.Args[1]
	if _, err := os.Stat(path); err == nil {

	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Invalid path")
		return
	} else {
		fmt.Println("Something is wrong with the file")
	}
	lib.ElaborateChat(path)
}
