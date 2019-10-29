package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1]
	file, err := os.Open(args) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := os.Stat(args)

	if err != nil {
		log.Fatal(err)
	}
	size := fileInfo.Size()

	data := make([]byte, size)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The contents of %q :\n%q\n", args, data[:count])

}
