package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// Guest is a structure modelling guest instance
type Guest struct {
	Id        int     `json:"user_id"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
	Name      string  `json:"name"`
}

// Guest.String() returns formatted string for pretty Guest visualisation
func (g *Guest) String() string {
	return fmt.Sprintf("Guest id: %d, name: %s, lat: %f, long: %f",
		g.Id, g.Name, g.Latitude, g.Longitude)
}

// Guest.isWithinDistance calculates whether a Guest object is within a given
// latitude and longitude point. Longitude and Latitude are passed in radians
// The method will return true if the guest is within a range, false otherwise.
// In case of problems, appropriate non-null error will be returned
func (g *Guest) isWithinDistance(longRad, latRad, distance float64) (bool, error) {

	if distance < 0.0 {
		return false, errors.New("Distance cannot be negative")
	}

	guestLat, guestLong, err := ConvertDegToRad(g.Latitude, g.Longitude)

	if err != nil {
		msg := fmt.Sprintf("Guest %d, error while converting Deg To Rad: %s\n",
			g.Id, err)
		return false, errors.New(msg)

	} else {

		longDelta := math.Abs(guestLong - longRad)

		angle := math.Acos(math.Sin(latRad)*math.Sin(guestLat) +
			math.Cos(latRad)*math.Cos(guestLat)*math.Cos(longDelta))

		return angle*EarthRadius <= distance, nil
	}
}

// Guests abstracts an iterable collection of *Guest objects
type Guests []*Guest

// Guests.Len() returns length of Guests objects
func (guests Guests) Len() int {
	return len(guests)
}

// Guests.Less() returns true if i'th element is smaller than j'th element of
// Guests's collection
func (guests Guests) Less(i, j int) bool {
	return guests[i].Id < guests[j].Id
}

// Guests.Swap() swaps i'th and j'th element in Guests collection
func (guests Guests) Swap(i, j int) {
	guest := guests[i]
	guests[i] = guests[j]
	guests[j] = guest
}

// Guests.FindGuestsWithinRange reurns a Guests collecion with all the Guest
// objects who are located within a given distance from the point depicted by
// longitude and latitude (both given in radians)
func (guests Guests) FindGuestsWithinRange(longitude, latitude, distance float64) Guests {
	guestsNearBy := make(Guests, 0, len(guests))

	var ok bool
	var err error

	for _, guest := range guests {
		ok, err = guest.isWithinDistance(
			longitude,
			latitude,
			distance)

		if err != nil {
			log.Println(err)
		} else if ok {
			guestsNearBy = append(guestsNearBy, guest)
		}
	}
	return guestsNearBy

}
