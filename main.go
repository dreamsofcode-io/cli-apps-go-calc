package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("missing subcommand")
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "add":
		addCmd(args)
	case "subtract":
		subtractCmd(args)
	default:
		log.Fatalln("invalid command")
	}
}

func addCmd(args []string) {
	if len(args) != 2 {
		log.Fatalln("incorrect arguments for add cmd")
	}

	const bitSize int = 64

	num1, err := strconv.ParseFloat(args[0], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	num2, err := strconv.ParseFloat(args[1], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	sum := num1 + num2

	fmt.Println(sum)
}

func subtractCmd(args []string) {
	if len(args) != 2 {
		log.Fatalln("incorrect arguments for add cmd")
	}

	const bitSize = 64

	num1, err := strconv.ParseFloat(args[0], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	num2, err := strconv.ParseFloat(args[1], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	sum := num1 - num2

	fmt.Println(sum)
}
