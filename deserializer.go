package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

// unMarshallGuest reads an array of bytes (which should be a JSON structure)
// and returns a *Guest object.
// If the array of bytes is malformed or JSON decoder cannot properly decode data
// and error will be returned
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

// ReadGuestsList accepts a filename with guests list and by reading file line
// by line tries to unmarshall and get a *Guest object instead.
// Effectively a Guests object is created and filled with *Guest objects and
// eventually returned.
// If the line cannot be properly decoded a log entry will be printed out
// and the line will be ommited.
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
