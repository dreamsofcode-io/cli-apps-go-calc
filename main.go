package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	log.SetFlags(0)

	var precision int
	flag.IntVar(&precision, "precision", -1, "The number of decimal places to print out")
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatalln("missing subcommand")
	}

	subcommand := flag.Arg(0)
	args := flag.Args()[1:]

	switch subcommand {
	case "add":
		addCmd(args, precision)
	case "subtract":
		subtractCmd(args, precision)
	default:
		log.Fatalln("invalid command")
	}
}

const bitSize = 64

func addCmd(args []string, precision int) {
	if len(args) != 2 {
		log.Fatalln("incorrect arguments for add cmd")
	}

	num1, err := strconv.ParseFloat(args[0], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	num2, err := strconv.ParseFloat(args[1], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	sum := num1 + num2

	printNumber(sum, precision)
}

func subtractCmd(args []string, precision int) {
	if len(args) != 2 {
		log.Fatalln("incorrect arguments for add cmd")
	}

	num1, err := strconv.ParseFloat(args[0], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	num2, err := strconv.ParseFloat(args[1], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	sum := num1 - num2

	printNumber(sum, precision)
}

func printNumber(num float64, precision int) {
	str := strconv.FormatFloat(num, 'f', precision, bitSize)
	fmt.Println(str)
}
