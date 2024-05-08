# Reservation System

<p align="left">
  <a href="https://github.com/hrishin/dockerfile-sources/actions"><img alt="GitHub Actions CI status" src="https://github.com/hrishin/dockerfile-sources/workflows/build-and-test/badge.svg"></a>
</p>

Is a CLI tool to mange the booking for flights reservations.

The tool is written in the [Golang](https://golang.org).

# Usage
### Docker
```bash
docker run hriships/reservation:v0.0.1
```

### Linux
```bash
export REPOSITORY_LIST_URL=https://gist.githubusercontent.com/jmelis/c60e61a893248244dc4fa12b946585c4/raw/25d39f67f2405330a6314cad64fac423a171162c/sources.txt 
````
```bash
./bin/reservation-linux-amd64 BOOK A0 2
```

### Mac
```bash
./bin/reservation-darwin-amd64 BOOK A0 2
```


## Extra Command
CLI provides following additional command
```
 VIEW        View flight booking
```

## Extra Options
Command provides following additional flags
```
--state-file string   directory path to store the booking state file
```

## Development
### Prerequisites
* Install the [Go](https://golang.org/doc/install) (1.21 >=)
* GNU [Make](https://www.gnu.org/software/make/)
* Docker (optional, but must be installed to build or run the container image)

### Building Binary
```
make all
```

### Building container image
```
make docker
```

## Testing

### Run Unit Tests
```
make unit-tests
```

### Run Integration Test
```
make integration-test
```

## Implementation Details
### Why CLI tool?
This preference allows users to execute the program on couple of environments such as Linux and Mac,

### System composition?
![alt text](/docs/flight-reservation.png "reservation system sources design")
Above diagram depicts objects composition that build up the CLI tool

## Design considerations
#### Extensibility
- As of now system supports `file based` booking system where all booking state stored in the file. However, system is extensible to support more providers
  using the `Storaable` interface, so database base providers are possible to support with the least possible changes.
- Most of the components are loosely coupled in abstracted way, so
  internal implementation can be altered without changing external interfaces; exception depends upon the use case.

## Assumptions, Limitations and Improvements
- Booking seats beyond row limits
    - As of user could request max 8 seats in one booking request with the given seat preference
    - However, system is extensible to support max number of seats with some changes to booking functionality

- Booking ID
  - As users requests booking, internally system generates the unique ID for the booking request. Id is represented as `int`.
  - For every booking request system increments the ID by 1, based on which processor system runs ID has finite range
  - In worst case if system run of the numbers for `int` space, then need to manually intervene in the statefile to reset the ID number
    or need to update the program to support better indexing scheme such as `RowID + random number`
