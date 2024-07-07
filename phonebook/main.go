package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//points := make(map[string]int)
//points["alice"] = 100
//points["bob"] = 200
//
//points2 := mao[stringint{
//	"Alice": 200,
//	"Bob":   100,
//}]

func main() {
	phoneBook := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Phonebook")
	fmt.Println("Available commands: add, get, delete, update, list, exit")

	for {
		fmt.Print("-> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			kv := strings.SplitN(parts[1], "=", 2)
			if len(kv) != 2 {
				fmt.Println("Invalid format. Use: add name=number")
				continue
			}
			name, number := kv[0], kv[1]
			phoneBook[name] = number
			fmt.Printf("Added/Updated: %s -> %s\n", name, number)

		case "get":
			name := parts[1]
			number, exists := phoneBook[name]
			if exists {
				fmt.Printf("Number for %s is %s\n", name, number)
			} else {
				fmt.Printf("No entry found for %s\n", name)
			}
		case "delete":
			name := parts[1]
			_, exists := phoneBook[name]
			if exists {
				delete(phoneBook, name)
				fmt.Printf("Deleted entry for %s\n", name)
			} else {
				fmt.Printf("No entry to delete for %s\n", name)
			}
		case "update":
			kv := strings.SplitN(parts[1], "=", 2)
			if len(kv) != 2 {
				fmt.Println("Invalid format. Use: add name=number")
				continue
			}
			name, number := kv[0], kv[1]
			_, exists := phoneBook[name]
			if exists {
				phoneBook[name] = number
				fmt.Printf("Updated entry for %s -> %s\n", name, number)
			} else {
				fmt.Println("No such entry to update.")
			}
		case "list":
			if len(phoneBook) == 0 {
				fmt.Println("Phone book is empty.")
			} else {
				for name, number := range phoneBook {
					fmt.Printf("%s -> %s\n", name, number)
				}
			}
		case "exit":
			fmt.Println("Exiting phonebook...")
			return
		default:
			fmt.Println("Unsupported command. Try 'add', 'get', 'delete', 'update', 'list', or 'exit'.")
		}
	}
}
