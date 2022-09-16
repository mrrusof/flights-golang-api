# flights-golang-api

A single endpoint Web API that summarizes a flight path as the start and end locations.

# Requirements

- Golang 1.8

# Getting started

1. Build: `make`
2. Run: `make run`
3. Make a sample request: `make sample_request`

# API Specification

The only endpoint of the Web API is `POST /summarize_flight_path`.
The endpoint summarizes a flight path consisting of one or more
legs.
The endpoint responds with a single leg corresponding to the start and
end locations of the path if the request body is a flight path.
Otherwise, the endpoint responds with HTTP
status code 400.

The request body is a JSON document with a structure corresponding to
the following Golang struct `FlightPath`.

```golang
type FlightPath struct {
  Path []FlightLeg
}

type FlightLeg struct {
  Origin string
  Destination string
}
```

Consider the following example request body.

```json
{
  "Path": [
    {
      "Origin": "ATL",
      "Destination": "EWR"
    },
    {
      "Origin": "SFO",
      "Destination": "ATL"
    }
  ]
}
```

A successful response is a JSON document with a structure
corresponding to the previous Golang struct `FlightLeg`.
Consider the following example response body that corresponds to the
previous request.

```json
{
  "Origin": "SFO",
  "Destination": "ATL"
}
```
