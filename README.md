# :hotel: Task Real Estate

[![Go Report Card](https://goreportcard.com/badge/github.com/baajarmeh/fiber-estate)](https://goreportcard.com/report/github.com/baajarmeh/fiber-estate)
[![CodeFactor](https://www.codefactor.io/repository/github/baajarmeh/fiber-estate/badge)](https://www.codefactor.io/repository/github/baajarmeh/fiber-estate)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/baajarmeh/fiber-estate)
![Lines of code](https://img.shields.io/tokei/lines/github/baajarmeh/fiber-estate)
![Github Repository Size](https://img.shields.io/github/repo-size/baajarmeh/fiber-estate)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/baajarmeh/fiber-estate)

Service for hotel room management and reservations.

# Run

```
make build
make run
```

if you are running the application for the first time you must apply migrations to the database:

```
make migrate_up
```

For migrations, use [golang-migrate/migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation).

# Unit tests

```
make run_test
```

# API

> 1. Request/response body - in JSON format.
> 2. In case of an error, the necessary HTTP code is returned, the body contains a description of the error (example: {"error": "something went wrong"}).

## POST /rooms/

Add a hotel room.

- Query body parameters:
  - description - text description,
  - price - price per night.
- Тело ответа:
  - room_id - hotel room id.

**Example**

Request:

```
curl -X POST localhost:2020/rooms/ \
-H "Content-Type: application/json" \
-d '{
	"description": "The Best Room",
	"price": 9000
}'
```

Response:

```
{
  "room_id": 144
}
```

## DELETE /rooms/:id

Deleting a hotel room.

- Query path parameters:
  - id - hotel room id.

**Example**

Request:

```
curl -X DELETE localhost:9000/rooms/144
```

## GET /rooms/

Get a list of hotel rooms.

- Query string parameters:
  - sort - the field by which the sort is performed:
    - id - by identifier (by date of addition),
    - price - at the price.
- Response body:
  - List of hotel rooms.

> 1. To sort in descending order, you must add a minus sign before the field value (-id or -price).
> 2. By default (if the sort parameter is empty or missing), sorting by id is ascending.

**Example**

Request:

```
curl -X GET localhost:9000/rooms/?sort=-price
```

Response:

```
[
  {
    "room_id": 2,
    "description": "description2",
    "price": 5000
  },
  {
    "room_id": 3,
    "description": "description3",
    "price": 3000
  },
  {
    "room_id": 1,
    "description": "description1",
    "price": 1000
  },
]
```

## POST /bookings/

Add a hotel room reservation.

- Query body parameters:
  - room_id - hotel room identifier,
  - date_start - the date of the beginning of the reservation,
  - date_end - the end date of the reservation.
- Response body:
  - booking_id - booking ID.

> Restrictions (from conditions): there is no check for the availability of a hotel room at the selected time.

**Example**

Request:

```
curl -X POST localhost:9000/bookings/ \
-H "Content-Type: application/json" \
-d '{
	"room_id": 144,
	"date_start": "2021-12-30",
	"date_end": "2022-01-02"
}'
```

Response:

```
{
  "booking_id": 121
}
```

## DELETE /bookings/:id

Delete a hotel room reservation.

- Query parameters:
  - id - booking id.

**Example**

Request:

```
curl -X DELETE localhost:9000/bookings/121
```

## GET /bookings/

Get a list of hotel room reservations.

- Query string parameters:
  - room_id - Hotel room IDENTIFIER.
- Response body:
  - list of reservations.

> The list is sorted by start date (date_start).

**Example**

Request:

```
curl -X GET localhost:9000/bookings/?room_id=144
```

Response:

```
[
  {
    "booking_id": 289,
    "date_start": "2021-01-04",
	  "date_end": "2021-01-08"
  },
  {
    "booking_id": 121,
    "date_start": "2021-12-30",
	  "date_end": "2022-01-02"
  },
  {
    "booking_id": 256,
    "date_start": "2022-03-01",
    "date_end": "2022-03-12"
  },
]
```

# Implementation

- Following the design of the REST JSON API.
- The "Pure Architecture" approach and the technique of addiction implementation.
- Work with the [fiber](https://github.com/gofiber/fiber).
- Working with the Postgres database using the [sqlx](https://github.com/jmoiron/sqlx) library and writing SQL queries.
- Application configuration - library [viper](https://github.com/spf13/viper).
- Graceful Shutdown implementation.
- Run from Docker.
- Unit testing with mocs - libraries [testify](https://github.com/stretchr/testify), [mock](https://github.com/golang/mock).
- Continuous integration, running tests in Travis CI, analysis and coverage in Scrutinizer CI.

**Project Structure**

```
.
├── pkg
│ ├── model       // basic structures
│ ├── handler     // request handlers
│ ├── service     // business logic
│ └── repository  // interaction with the back-up
├── cmd           // application entry point
├── scripts       // SQL files with migrations
└── configs       // configuration files
```
