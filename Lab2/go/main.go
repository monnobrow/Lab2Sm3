package main

import (
	"lab2ann/tasks"
	"os"
)

func main() { os.Exit(tasks.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)) }
