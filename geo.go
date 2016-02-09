package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	DegToRad           = math.Pi / 180
	DublinLatitudeDeg  = 53.3381985
	DublinLongitudeDeg = -6.2592576
	DublinLatitudeRad  = DublinLatitudeDeg * DegToRad
	DublinLongitudeRad = DublinLongitudeDeg * DegToRad
	EarthRadius        = 6371.0 // in km
)

func ConvertDegToRad(lat, long float64) (float64, float64, error) {

	if lat < -90 || lat > 90 {
		msg := fmt.Sprintf("Latitude must be withing range [-90,90], got %d", lat)
		return 0.0, 0.0, errors.New(msg)
	}

	if long < -180 || long > 180 {
		msg := fmt.Sprintf("Longitude must be withing range [-180,180], got %d", long)
		return 0.0, 0.0, errors.New(msg)
	}

	return lat * DegToRad, long * DegToRad, nil
}
