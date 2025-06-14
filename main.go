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

	// Homework Part 1: Global rounding flags
	var round bool
	flag.BoolVar(&round, "round", false, "Round result to nearest integer")

	var floor bool
	flag.BoolVar(&floor, "floor", false, "Round result down to nearest integer")

	var ceil bool
	flag.BoolVar(&ceil, "ceil", false, "Round result up to nearest integer")

	flag.Parse()

	// Validate mutually exclusive rounding flags
	flagCount := 0
	if round {
		flagCount++
	}
	if floor {
		flagCount++
	}
	if ceil {
		flagCount++
	}
	if flagCount > 1 {
		log.Fatalln("Cannot use multiple rounding flags together")
	}

	if len(flag.Args()) < 1 {
		log.Fatalln("missing subcommand")
	}

	subcommand := flag.Arg(0)
	args := flag.Args()[1:]

	// Create rounding config to pass to commands
	roundingConfig := RoundingConfig{
		Round: round,
		Floor: floor,
		Ceil:  ceil,
	}

	switch subcommand {
	case "add":
		addCmd(args, precision, roundingConfig)
	case "subtract":
		subtractCmd(args, precision, roundingConfig)
	case "multiply":
		multiplyCmd(args, precision, roundingConfig)
	case "divide":
		divideCmd(args, precision, roundingConfig)
	default:
		log.Fatalln("invalid command")
	}
}

type RoundingConfig struct {
	Round bool
	Floor bool
	Ceil  bool
}

const bitSize = 64

func addCmd(args []string, precision int, rounding RoundingConfig) {
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

	printNumber(sum, precision, rounding)
}

func subtractCmd(args []string, precision int, rounding RoundingConfig) {
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

	printNumber(result, precision, rounding)
}

func multiplyCmd(args []string, precision int, rounding RoundingConfig) {
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

	printNumber(result, precision, rounding)
}

func divideCmd(args []string, precision int, rounding RoundingConfig) {
	// Homework Part 2: Remainder flag for divide command
	var isRemainder bool

	flagSet := flag.NewFlagSet("divide", flag.ExitOnError)
	flagSet.BoolVar(
		&isRemainder,
		"remainder",
		false,
		"Return the integer remainder instead of division result",
	)
	flagSet.Parse(args)

	args = flagSet.Args()

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

	var result float64

	if isRemainder {
		result = math.Mod(num1, num2)
	} else {
		result = num1 / num2
	}

	printNumber(result, precision, rounding)
}

func printNumber(num float64, precision int, rounding RoundingConfig) {
	// Apply rounding based on global flags
	if rounding.Round {
		num = math.Round(num)
	} else if rounding.Floor {
		num = math.Floor(num)
	} else if rounding.Ceil {
		num = math.Ceil(num)
	}

	str := strconv.FormatFloat(num, 'f', precision, bitSize)
	fmt.Println(str)
}

