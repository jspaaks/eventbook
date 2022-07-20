# Simple REST API web server written in Go

Basic CRUD'ing of the elements of an array via REST requests to a web server. Demonstration of:

- Closures
- Generics
- Pointers
- REST API with POST/GET/PUT/DELETE requests, see below

## Endpoints

`POST /api/events`

Append one event to the list of events. Events are layed out as defined in `src/types/types.go`.

`GET /api/events`

Get all events.

`GET /api/event/{index:[0-9]+}`

Get one event by its index.

`PUT /api/event/{index:[0-9]+}`

Update one event by its index. Events are layed out as defined in `src/types/types.go`.

`DELETE /api/event/{index:[0-9]+}`

Delete one event by its index.

## Prerequisites

- Go >= 1.18 (because of generics)

## Starting the server

```shell
git clone https://github.com/jspaaks/eventbook .
make
```

The webserver will be on http://localhost:8080
