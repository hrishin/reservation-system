# Reservation System

<p align="left">
  <a href="https://github.com/hrishin/reservation-system/actions"><img alt="GitHub Actions CI status" src="https://github.com/hrishin/reservation-system/workflows/build-and-test/badge.svg"></a>
</p>

Is a CLI tool to mange the booking for flights reservations.

The tool is written in the [Golang](https://golang.org).

# Usage
### Docker
```bash
mkdir -p $(pwd)/datadir && docker run -v $(pwd)/datadir/:/tmp/data hriships/reservation:v0.0.1 BOOK A0 1 --state-file /tmp/data
```

### Binary

Download the right binary from https://github.com/hrishin/reservation-system/releases

```bash
rservation A0 1 --state-file /tmp/data
```

## Extra Command
CLI provides the following additional command
```
 VIEW        View flight booking
```

## Extra Options
The command provides the following additional flags
```
--state-file string   directory path to store the booking state file
-v, --verbose         enables verbose logging


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
This preference allows users to execute the program on a couple of environments such as Linux and Mac,

### System composition?
![alt text](/docs/flight-reservation.png "reservation system sources design")
The above diagram depicts object composition that builds up the CLI tool

## Design considerations
#### Extensibility
- As of now system supports `file based` booking system where all booking state are stored in the file. However, the system is extensible to support more providers
  using the `Storaable` interface, so database base providers are possible to support with the least possible changes.
- Most of the components are loosely coupled in abstract way, so
  internal implementation can be altered without changing external interfaces; exception depends upon the use case.

## Assumptions, Limitations and Improvements
- Booking seats beyond row limits
    - As of user could request max 8 seats in one booking request with the given seat preference
    - However, the system is extensible enough to support max number of seats with some changes to the booking functionality

- Booking ID
  - As users request a booking, internal system generates the unique ID for the booking request. ID is represented as `int`.
  - For every booking request system increments the ID by 1, based on which processor system runs ID has finite range
  - In the worst case, if the system runs off the numbers for `int` space, then need to manually intervene in the statefile to reset the ID number
    or need to update the program to support better indexing scheme such as `RowID + random number`

- Integration tests
  - Some of the code is redundant for a reason, as tests should be no-brainer and easy to read and interpret 
  - Additional scenarios could be covered such as checking view command, verbose logs 
