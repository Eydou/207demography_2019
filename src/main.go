//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// main
//

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	functions "./functions"
)

func help() {
	fmt.Printf("USAGE\n   ./207demography code [...]\n")
	fmt.Printf("\nDESCRIPTION\n")
	fmt.Printf("   code\t    country code\n")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			help()
		}
	}
	if _, err := functions.ErrorArgs(args); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mX\033[0m Error: %s\n", err)
		os.Exit(84)
	}
	data, err := ioutil.ReadFile("207demography_data.csv")
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(84)
	}
	os.Exit(functions.MathParse(string(data), args))
}
