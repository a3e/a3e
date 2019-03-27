package main

import "os"

func main() {
	cmd := Root()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
