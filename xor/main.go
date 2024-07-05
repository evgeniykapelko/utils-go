package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()

	if len(*secretKey) == 0 {
		fmt.Fprintf(os.Stderr, "You must provide a secret key.\r\n")
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		plaintext := getUserInput("Enter your text to cipher: ")
		fmt.Println("Plaintext:", plaintext)

		cipheredText, err := cipherer.Cipher(plaintext, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encrypting text: %v\r\n", err)
			os.Exit(1)
		}

		fmt.Println("Ciphered text:", cipheredText)
	case "decipher":

		cipheredText := getUserInput("Enter your text to decipher: ")
		decipheredText, err := cipherer.Decipher(cipheredText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decrypting text: %v\r\n", err)
			os.Exit(1)
		}

		fmt.Println("Deciphered text:", decipheredText)
	default:
		fmt.Println("Unsupported mode")
		os.Exit(1)
	}
}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			continue
		}

		return strings.TrimRight(result, "\r\n")
	}
}
