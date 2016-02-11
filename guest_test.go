package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type Location struct {
	Name string
	Lat  float64
	Long float64
}

func TestGuestDistanceCalculations(t *testing.T) {
	warsaw := Guest{
		Id:        0,
		Name:      "Warsaw",
		Latitude:  52.14,
		Longitude: 21.1,
	}

	cities := []Location{
		Location{"Otwock", 52.249 * DegToRad, 21.1448 * DegToRad},
		Location{"Minsk", 52.11 * DegToRad, 21.34 * DegToRad},
		Location{"Pruszkow", 52.10 * DegToRad, 20.48 * DegToRad},
	}

	distance := 50.0 // km

	for _, city := range cities {
		if ok, err := warsaw.isWithinDistance(city.Long, city.Lat, distance); err == nil {
			if ok == false {
				t.Errorf("City %s is within a range of %f kilometers, yet it was incorrectly calculated.", city.Name, distance)
			}
		} else {
			t.Errorf("Got error %s\n", err)
		}
	}
}

func TestGuestDistanceCalculationsFarAway(t *testing.T) {
	warsaw := Guest{
		Id:        0,
		Name:      "Warsaw",
		Latitude:  52.14,
		Longitude: 21.1,
	}

	cities := []Location{
		Location{"Geneva", 46.12 * DegToRad, 6.09 * DegToRad},
		Location{"Zurich", 47.22 * DegToRad, 8.33 * DegToRad},
		Location{"New York", 43.0 * DegToRad, -75.0 * DegToRad},
		Location{"Dublin", 53.2052 * DegToRad, -6.1535 * DegToRad},
	}

	distance := 500.0 // km

	for _, city := range cities {
		if ok, err := warsaw.isWithinDistance(city.Long, city.Lat, distance); err == nil {
			if ok {
				t.Errorf("City %s isn't within a range of %f km, yet the calculations claim it is",
					city.Name, distance)
			}
		} else {
			t.Errorf("Got error: %s\n", err)
		}
	}
}

func TestGuestDistanceCalculationsInvalidDistance(t *testing.T) {
	warsaw := Guest{
		Id:        0,
		Name:      "Warsaw",
		Latitude:  52.14,
		Longitude: 21.1,
	}

	if _, err := warsaw.isWithinDistance(0.0, 0.0, -1); err == nil {
		t.Error("Distance was negative, expected false, got true")
	}
}

func TestGuestsDistanceCalculations(t *testing.T) {
	guests := Guests{
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
		&Guest{
			Id:        2,
			Name:      "New York",
			Latitude:  43.0,
			Longitude: -75.0,
		},
		&Guest{
			Id:        3,
			Name:      "Dublin",
			Latitude:  53.2052,
			Longitude: -6.1535,
		},
	}

	WarsawLatRad := 52.14 * DegToRad
	WarsawLongRad := 21.1 * DegToRad

	guestsNearBy := guests.FindGuestsWithinRange(WarsawLongRad, WarsawLatRad,
		1000.0)

	if len(guestsNearBy) != 0 {
		fmt.Printf("None of the cities is withing distance %f km from Warsaw, yet found some:", distance)
		for _, g := range guestsNearBy {
			fmt.Println(g)
		}
		t.Error("List should be empty")
	}
}

func TestGuestsDistancePositiveResults(t *testing.T) {
	guests := Guests{
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
		&Guest{
			Id:        2,
			Name:      "New York",
			Latitude:  43.0,
			Longitude: -75.0,
		},
		&Guest{
			Id:        3,
			Name:      "Dublin",
			Latitude:  53.2052,
			Longitude: -6.1535,
		},
	}

	WarsawLatRad := 52.14 * DegToRad
	WarsawLongRad := 21.1 * DegToRad

	guestsNearBy := guests.FindGuestsWithinRange(WarsawLongRad, WarsawLatRad,
		1200.0)

	if len(guestsNearBy) != 1 {
		t.Error("Expected only one guest to be withing a range")
	}

	guest := guestsNearBy[0]
	if reflect.DeepEqual(guests[1], guest) == false {
		t.Errorf("Expected different city\nGot: %s\nExpected: %s\n",
			guest, guests[1])
	}
}

// test sorting methods

func TestGuestsLen(t *testing.T) {
	guests := Guests{
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
		&Guest{
			Id:        2,
			Name:      "New York",
			Latitude:  43.0,
			Longitude: -75.0,
		},
		&Guest{
			Id:        3,
			Name:      "Dublin",
			Latitude:  53.2052,
			Longitude: -6.1535,
		},
	}

	if guests.Len() != 4 {
		t.Errorf("Expected length to be 4, got %d instead\n", guests.Len())
	}
}

func TestGuestsLess(t *testing.T) {
	guests := Guests{
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
		&Guest{
			Id:        2,
			Name:      "New York",
			Latitude:  43.0,
			Longitude: -75.0,
		},
		&Guest{
			Id:        3,
			Name:      "Dublin",
			Latitude:  53.2052,
			Longitude: -6.1535,
		},
	}

	if guests.Less(0, 1) == false {
		t.Error("Expected asserted elements to indeed be less")
	}

	if guests.Less(1, 0) == true {
		t.Error("Expected asserted elements NOT to be less")
	}
}

func TestGuestsSwap(t *testing.T) {
	guests := Guests{
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
	}

	guests.Swap(0, 1)

	exp_guests := Guests{
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
	}

	if reflect.DeepEqual(exp_guests, guests) == false {
		t.Errorf("Expected Guests.Swap() to swap elements\n")
	}
}

func TestGuestsSortOnId(t *testing.T) {
	guests := Guests{
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
		&Guest{
			Id:        3,
			Name:      "Dublin",
			Latitude:  53.2052,
			Longitude: -6.1535,
		},
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
		&Guest{
			Id:        2,
			Name:      "New York",
			Latitude:  43.0,
			Longitude: -75.0,
		},
	}

	exp_guests := Guests{
		&Guest{
			Id:        0,
			Name:      "Geneva",
			Latitude:  46.12,
			Longitude: 6.09,
		},
		&Guest{
			Id:        1,
			Name:      "Zurich",
			Latitude:  47.22,
			Longitude: 8.33,
		},
		&Guest{
			Id:        2,
			Name:      "New York",
			Latitude:  43.0,
			Longitude: -75.0,
		},
		&Guest{
			Id:        3,
			Name:      "Dublin",
			Latitude:  53.2052,
			Longitude: -6.1535,
		},
	}

	sort.Sort(guests)

	if reflect.DeepEqual(exp_guests, guests) == false {
		fmt.Println("Expected guests list to be sorted on Id", "Got:")
		for _, guest := range guests {
			fmt.Println(guest)
		}

		fmt.Println("Expected:")
		for _, guest := range exp_guests {
			fmt.Println(guest)
		}
		t.Error("Guests list is not properly sorted")
	}
}
