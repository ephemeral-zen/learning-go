package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path,
		"path",
		".",
		"Path to the json file")
	flag.Parse()
	bookworms, err := loadBookworms(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr,
			"failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	//fmt.Println(bookworms)
	//findCommonBooks(bookworms)
	fmt.Print(recommendOtherBooks(bookworms))
}
