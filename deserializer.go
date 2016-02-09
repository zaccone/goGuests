package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func unMarshallGuest(data []byte) (*Guest, error) {
	var guest Guest
	err := json.Unmarshal(data, &guest)
	if err != nil {
		log.Printf("Error while unmarshalling line '%s' (will skip): %s\n",
			data, err)
		return nil, err
	} else {
		return &guest, nil
	}

}

func ReadGuestsList(input string) (Guests, error) {

	guests := make(Guests, 0, 0)

	if inputHandler, err := os.Open(input); err != nil {
		return nil, err
	} else {
		defer inputHandler.Close()
		scanner := bufio.NewScanner(inputHandler)

		for scanner.Scan() {
			rawInput := scanner.Bytes()
			if guest, err := unMarshallGuest(rawInput); err == nil {
				guests = append(guests, guest)
			}
		}
	}

	return guests, nil

}
