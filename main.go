package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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
	case "multiply":
		multiplyCmd(args, precision)
	case "divide":
		divideCmd(args, precision)
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
	var isAbsolute bool

	flagSet := flag.NewFlagSet("subtract", flag.ExitOnError)
	flagSet.BoolVar(
		&isAbsolute,
		"abs",
		false,
		"Determines whether or not to print out the absolute value",
	)
	flagSet.Parse(args)

	args = flagSet.Args()

	if len(args) != 2 {
		log.Fatalln("incorrect arguments for subtract cmd")
	}

	num1, err := strconv.ParseFloat(args[0], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	num2, err := strconv.ParseFloat(args[1], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	result := num1 - num2

	if isAbsolute {
		result = math.Abs(result)
	}

	printNumber(result, precision)
}

func multiplyCmd(args []string, precision int) {
	if len(args) != 2 {
		log.Fatalln("incorrect arguments for multiply cmd")
	}

	num1, err := strconv.ParseFloat(args[0], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	num2, err := strconv.ParseFloat(args[1], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	result := num1 * num2

	printNumber(result, precision)
}

func divideCmd(args []string, precision int) {
	if len(args) != 2 {
		log.Fatalln("incorrect arguments for divide cmd")
	}

	num1, err := strconv.ParseFloat(args[0], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	num2, err := strconv.ParseFloat(args[1], bitSize)
	if err != nil {
		log.Fatalln("invalid number:", err)
	}

	// Guard against division by zero
	if num2 == 0 {
		log.Fatalln("cannot divide by zero")
	}

	result := num1 / num2

	printNumber(result, precision)
}

func printNumber(num float64, precision int) {
	str := strconv.FormatFloat(num, 'f', precision, bitSize)
	fmt.Println(str)
}

