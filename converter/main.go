package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("Enter hex number or 'stop' to exit:")

		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) == "stop" {
			break
		}

		i := new(big.Int)

		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("Invalid input")
			continue
		}

		var base int
		fmt.Println("Enter the base to convert to (2, 8, 10, 16):")
		fmt.Scanln(&base)

		switch base {
		case 2:
			fmt.Println("Binary:", i.Text(2))
		case 8:
			fmt.Println("Octal:", i.Text(8))
		case 10:
			fmt.Println("Decimal:", i.Text(10))
		case 16:
			fmt.Println("Hexadecimal:", i.Text(16))
		default:
			fmt.Println("Invalid base. Please choose from 2, 8, 10, 16.")
			os.Exit(1)
		}
	}
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
