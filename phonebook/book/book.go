package book

import (
	"errors"
	"fmt"
	"time"
)

type PhoneBook map[string]PhoneNumber

type PhoneNumber struct {
	Number        string
	LastUpdatedAt int64
}

func (book *PhoneBook) Add(name string, phoneNum string) error {
	if _, ok := (*book)[phoneNum]; ok {
		return errors.New("phone number already exists")
	}

	//n := new(PhoneNumber)
	//n.Number = number
	//n.LastUpdatedAt = time.Now().Unix()

	(*book)[name] = PhoneNumber{
		Number:        phoneNum,
		LastUpdatedAt: time.Now().Unix(),
	}

	return nil
}

func (book *PhoneBook) Get(number string) (PhoneNumber, error) {
	if numberData, ok := (*book)[number]; ok {
		return numberData, nil
	}

	return PhoneNumber{}, fmt.Errorf("No entry found for phone number %s", number)
}

func (book *PhoneBook) Update(name string, number string) error {
	if _, ok := (*book)[name]; ok {
		(*book)[name] = PhoneNumber{
			Number:        number,
			LastUpdatedAt: time.Now().Unix(),
		}
	} else {
		return fmt.Errorf("No entry found for phone number %s", number)
	}

	return nil
}

func (book *PhoneBook) Delete(name string) error {
	if _, ok := (*book)[name]; ok {
		delete(*book, name)
	} else {
		return fmt.Errorf("No entry found for name %s", name)
	}

	return nil
}
