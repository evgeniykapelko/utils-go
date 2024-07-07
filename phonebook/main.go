package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"phonebook/book"
	"phonebook/logger"
	"strings"
	"time"
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
	phoneBook := make(book.PhoneBook)

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
		args := parts[1:]

		switch command {
		case "add":
			handleCommand(doAdd, args, phoneBook)
		case "get":
			handleCommand(doGet, args, phoneBook)

		case "delete":
			handleCommand(doDelete, args, phoneBook)

		case "update":
			handleCommand(doUpdate, args, phoneBook)

		case "list":
			handleCommand(doList, args, phoneBook)

		case "exit":
			logger.Info("Exiting phonebook, bye!")
			return
		default:
			logger.Warn(errors.New("Unsupported command. Try 'add', 'get', 'delete', 'update', 'list', or 'exit'."))
		}
	}
}

func handleCommand(
	cmd func([]string, book.PhoneBook) error,
	args []string,
	phoneBook book.PhoneBook,
) {
	if err := cmd(args, phoneBook); err != nil {
		logger.Warn(err, "command failed")
	}
}

func doAdd(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("Not enough arguments")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("Wrong number of arguments")
	}

	name, number := kv[0], kv[1]
	err := phoneBook.Add(name, number)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Added phone book with name %s and number %s", name, number))

	return nil
}

func doGet(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("Not enough arguments")
	}

	name := args[0]

	numberData, err := phoneBook.Get(name)
	if err != nil {
		return err
	}

	unixUpdatedAt := time.Unix(numberData.LastUpdatedAt, 0)

	logger.Info(
		fmt.Sprintf(
			"Number for %s is %s (last updated at %s)",
			name,
			numberData.LastUpdatedAt,
			unixUpdatedAt.Format("2006-01-02 15:04:05"),
		),
	)

	return nil
}

func doList(_ []string, phoneBook book.PhoneBook) error {
	if len(phoneBook) == 0 {
		return errors.New("No phone book found")
	} else {
		results := ""

		for name, number := range phoneBook {
			results += fmt.Sprintf("\n%s -> %s", name, number.Number)
		}

		logger.Info(results)
	}

	return nil
}

func doUpdate(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("Not enough arguments")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("Wrong number of arguments")
	}

	name, number := kv[0], kv[1]
	err := phoneBook.Update(name, number)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Updated an entry: %s -> %s\n", name, number))

	return nil
}

func doDelete(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("Not enough arguments")
	}

	name := args[0]

	err := phoneBook.Delete(name)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Deleted entry for %s\n", name))

	return nil
}
