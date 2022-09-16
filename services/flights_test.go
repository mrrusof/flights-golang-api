package services

import (
  "fmt"
  "math/rand"
  "testing"
)

func TestSummarize(t *testing.T) {
  cases := []struct {
    input [][]string
    expected []string
  }{

    { // Single leg
      [][]string{{"SFO", "ATL"}},
      []string{"SFO", "ATL"},
    },

    { // No legs
      [][]string{},
      []string{"", ""},
    },

    { // Two legs in order
      [][]string{{"SFO", "ATL"}, {"ATL", "EWR"}},
      []string{"SFO", "EWR"},
    },

    { // Two legs reversed
      [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}},
      []string{"SFO", "EWR"},
    },

    { // Two flight origins
      [][]string{{"ATL", "DFW"}, {"SFO", "DFW"}},
      []string{"", ""},
    },

    { // Two flight destinations
      [][]string{{"DFW", "ATL"}, {"DFW", "SFO"}},
      []string{"", ""},
    },

    { // 1M legs
      generate1MRandomLegs(),
      []string{"0", "1000000"},
    },
  }

  for _, raw := range cases {
    input := sliceToFlightPath(raw.input)
    expected := pairToFlightLeg(raw.expected)
    actual := Summarize(input)
    if actual != expected {
      t.Errorf("Summarize(%.50s) == %q, expected %q", fmt.Sprint(input), actual, expected)
    }
  }
}

func pairToFlightLeg(p []string) FlightLeg {
  return FlightLeg{ Origin: p[0], Destination: p[1] }
}

func sliceToFlightPath(s [][]string) FlightPath {
  flight := FlightPath{}
  flight.Path = make([]FlightLeg, len(s))
  for i, p := range s {
    flight.Path[i] = pairToFlightLeg(p)
  }
  return flight
}

func generate1MRandomLegs() [][]string {
  legs := make([][]string, 1000000)

  for i := range legs {
    legs[i] = []string{fmt.Sprint(i), fmt.Sprint(i + 1)}
  }

  for i := range legs {
    j := rand.Intn(i + 1)
    legs[i], legs[j] = legs[j], legs[i]
  }

  return legs
}
