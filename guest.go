package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type Guest struct {
	Id        int     `json:"user_id"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
	Name      string  `json:"name"`
}

func (g *Guest) String() string {
	return fmt.Sprintf("Guest id: %d, name: %s, lat: %f, long: %f",
		g.Id, g.Name, g.Latitude, g.Longitude)
}

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

type Guests []*Guest

func (guests Guests) Len() int {
	return len(guests)
}

func (guests Guests) Less(i, j int) bool {
	return guests[i].Id < guests[j].Id
}

func (guests Guests) Swap(i, j int) {
	guest := guests[i]
	guests[i] = guests[j]
	guests[j] = guest
}

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
