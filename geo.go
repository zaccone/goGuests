package main

import (
	"math"
)

const (
	DegToRad           = math.Pi / 180
	DublinLatitudeRad  = 53.3381985 * DegToRad
	DublinLongitudeRad = -6.2592576 * DegToRad
	EarthRadius        = 6371.0 // in km
)

func ConvertDegToRad(lat, long float64) (float64, float64) {
	return lat * DegToRad, long * DegToRad
}
