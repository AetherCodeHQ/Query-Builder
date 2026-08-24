package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: query-builder <table> [cols] [--where k=v] [--order col] [--limit N]")
		os.Exit(1)
	}
	table := os.Args[1]
	cols := "*"
	where, order, limit := "", "", ""
	args := os.Args[2:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--where":
			if i+1 < len(args) {
				parts := strings.SplitN(args[i+1], "=", 2)
				if len(parts) == 2 {
					where = fmt.Sprintf(" WHERE %s = '%s'", parts[0], parts[1])
				}
				i++
			}
		case "--order":
			if i+1 < len(args) {
				order = " ORDER BY " + args[i+1]
				i++
			}
		case "--limit":
			if i+1 < len(args) {
				limit = " LIMIT " + args[i+1]
				i++
			}
		default:
			cols = args[i]
		}
	}
	fmt.Printf("SELECT %s FROM %s%s%s%s;\n", cols, table, where, order, limit)
}
