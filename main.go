package main

import (
	"fmt"
	"os"
)

// query_builder - SQL query builder
func query_builder(path string) {
	fmt.Println("========================================")
	fmt.Println("  Query-Builder")
	fmt.Println("  SQL query builder")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	query_builder(path)
}
