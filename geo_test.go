package main

import (
	"math"
	"testing"
)

func TestConvertingDegtoRad(t *testing.T) {

	type Pair struct {
		Lat  float64
		Long float64
	}

	degrees := []Pair{
		Pair{-90, -180},
		Pair{0, 0},
		Pair{90, 180},
		Pair{45, 45},
	}

	radians := []Pair{
		Pair{-math.Pi / 2, -math.Pi},
		Pair{0, 0},
		Pair{math.Pi / 2, math.Pi},
		Pair{math.Pi / 4, math.Pi / 4},
	}

	for i := 0; i < len(degrees); i++ {
		latRad, longRad, err := ConvertDegToRad(degrees[i].Lat,
			degrees[i].Long)

		if err != nil {
			t.Errorf("Didn't expect any error, yet got one %s\n", err)
		}

		if latRad != radians[i].Lat {
			t.Errorf("Latitude conversion mismatch, expected %f, got %f\n",
				radians[i].Lat, latRad)
		}

		if longRad != radians[i].Long {
			t.Errorf("Longitude conversion mismatch, expected %f, got %f\n",
				radians[i].Long, longRad)
		}
	}
}

func TestConvertingDegtoRadInvalidInput(t *testing.T) {

	type Pair struct {
		Lat  float64
		Long float64
	}

	degrees := []Pair{
		Pair{-783, -180},
		Pair{200, 200.8},
		Pair{-181.8, 270.0},
		Pair{45.5, 450.5},
	}

	for _, test := range degrees {
		lat, long, err := ConvertDegToRad(test.Lat,
			test.Long)
		if err == nil {
			t.Errorf("Expected error to be not null, testcase: lat: %f, long: %f\n",
				test.Lat, test.Long)
		}

		if lat != 0.0 || long != 0.0 {
			t.Errorf("Expected latitude and longitude to be set to 0.0, got %f, %f instead\n",
				lat, long)
		}
	}
}
