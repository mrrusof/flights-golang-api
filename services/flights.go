package services

type FlightPath struct {
  Path []FlightLeg
}

type FlightLeg struct {
  Origin string
  Destination string
}

func Summarize(flight FlightPath) FlightLeg {
  if len(flight.Path) == 0 {
    return FlightLeg{}
  }

  origins := make(map[string]bool)
  destinations := make(map[string]bool)

  for _, leg := range flight.Path {
    origins[leg.Origin] = true
    destinations[leg.Destination] = true
  }

  for _, leg := range flight.Path {
    origins[leg.Destination] = false
    destinations[leg.Origin] = false
  }

  leg := FlightLeg{}

  for loc, isOrigin := range origins {
    if isOrigin {
      if leg.Origin == "" {
        leg.Origin = loc
      } else {
        return FlightLeg{}
      }
    }
  }

  for loc, isDestination := range destinations {
    if isDestination {
      if leg.Destination == "" {
        leg.Destination = loc
      } else {
        return FlightLeg{}
      }
    }
  }

  return leg
}
