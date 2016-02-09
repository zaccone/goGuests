package main

import (
	"reflect"
	"testing"
)

func TestUnMarshallGuest(t *testing.T) {
	line := `{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`

	guest, err := unMarshallGuest([]byte(line))
	if err != nil {
		t.Error("Error while umarshalling input data: ", err)
	}

	refGuest := &Guest{
		Id:        12,
		Latitude:  52.986375,
		Longitude: -6.043701,
		Name:      "Christina McArdle",
	}

	if reflect.DeepEqual(refGuest, guest) == false {
		t.Errorf("Guests mismatch:\nExp %s\nGot: %s", refGuest, guest)
	}
}

func TestUnMarshallGuestInvalidInput(t *testing.T) {
	line := `invalid input`
	guest, err := unMarshallGuest([]byte(line))
	if err == nil || guest != nil {
		t.Error(
			"Error while umarshalling input data, expected function to fail")
	}
}

func TestReadingGuestListFromFile(t *testing.T) {
	fileName := "fixtures/guests_single_line.txt"
	guests, err := ReadGuestsList(fileName)
	if err != nil {
		t.Error("Error while unmarshalling guests from a file", err)
	}

	if len(guests) != 1 {
		t.Errorf("Expected to get a single item list, got %d instead\n",
			len(guests))
	}

	refGuest := &Guest{
		Id:        12,
		Latitude:  52.986375,
		Longitude: -6.043701,
		Name:      "Christina McArdle",
	}

	if reflect.DeepEqual(refGuest, guests[0]) == false {
		t.Errorf("Guests mismatch:\nExp %s\nGot: %s", refGuest, guests[0])
	}
}

func TestReadingGuestListFromFileMulti(t *testing.T) {
	fileName := "fixtures/guests.txt"
	expGuestLength := 32

	guests, err := ReadGuestsList(fileName)
	if err != nil || guests == nil {
		t.Error("Error while unmarshalling guests from a file", err)
	}

	if len(guests) != expGuestLength {
		t.Errorf("Expected to get a list with %d elements, got %d instead\n",
			expGuestLength, len(guests))
	}
}

func TestReadingGuestFileWrongPath(t *testing.T) {
	fileName := "fixtures/idontexist"
	guests, err := ReadGuestsList(fileName)

	if err == nil {
		t.Error("Expected bad input file to be handled")
	}

	if len(guests) != 0 {
		t.Error("Expected guests list to be empty")
	}
}

func TestReadingGuestFileInvalidInput(t *testing.T) {
	fileName := "fixtures/guests_invalid_input.txt"
	guests, err := ReadGuestsList(fileName)

	if err != nil {
		t.Error("ReadGuestsList should not set error in this case (invalid input)")
	}

	if len(guests) != 0 {
		t.Errorf("Expected guest list to be empty, instead it's %d\n",
			len(guests))
	}

}
