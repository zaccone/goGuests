package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"sort"
)

const defaultDistance = 100.0 // in km

var fileName string
var distance float64

// parse parses available options and guards invalid options.
// As this is usually beginning of whole application, all failures will
// result in error message being displayed and exiting program.
func parse() error {

	flag.StringVar(&fileName, "f", "", "Path of a file with list of guests")
	flag.Float64Var(&distance, "d", defaultDistance,
		"Distance in km, defaults to 100 km")
	flag.Parse()

	if fileName == "" {
		return errors.New("File path must not be empty")
	}

	if distance < 0 {
		return errors.New("Distance cannot be negative")
	}

	return nil
}

func main() {

	if err := parse(); err != nil {
		log.Fatal(err)
	}

	if guests, err := ReadGuestsList(fileName); err != nil {
		log.Fatal(err)
	} else {
		inviteList := guests.FindGuestsWithinRange(
			DublinLongitudeRad, DublinLatitudeRad, distance)
		sort.Sort(inviteList)
		for i, guest := range inviteList {
			fmt.Println(i, guest)
		}
	}

}
