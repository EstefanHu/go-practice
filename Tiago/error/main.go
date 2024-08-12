package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	user, err := getUser(42)
	if err != nil {
		log.Printf("main error: %v", err.Error())
		return
	}

	err = doWork(user)
	if err != nil {
		return
	}

	err = doWork(user)
	if err != nil {
		return
	}
}

func doWork(user string) error {
	fmt.Printf("Doing work for %s", user)
	return nil
}

func getUser(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("id cannot be negative")
	}

	return "foo", nil
}
