## Tutorials 🎥

- [Codemy.com](https://youtu.be/YAfRxJBELv0)

- [ProgrammingKnolodge - Go Tutorial (Golang) 16 - Structures in Go (structs)](https://youtu.be/t24UEghfyvo)

- [ProgrammingKnolodge - Go Tutorial (Golang) 17 - Composition in Golang](https://youtu.be/-FN7h2xM0W8)

- [Creating a Simple RESTful API in Go - Part 1](https://youtu.be/W5b64DXeP0o)

- [Creating A RESTful API in Go - Part 2](https://youtu.be/YMQUQ6XQgz8)

- [Consume a REST API in Go | Convert JSON to a Struct](https://youtu.be/aYk8XAKxhxU)

- [Go API Tutorial - Make An API With Go](https://youtu.be/bj77B59nkTQ)

- [How to create JSON data in golang](https://youtu.be/SZ5xZ9OTeEI)

- [How to consume JSON data in golang](https://youtu.be/a96veXdifys)

## Some tools 🛠️

- [Transform.tool](https://transform.tools/json-to-go)

- [Postman](https://www.postman.com/)

## Previous project 🗃️

- [Medina](https://github.com/Medina7276/groupie-tracker_Golang)

## Structs model 🔩

```bash

// Artist struct represents information about a band or artist.
type Artist struct {
	Name         string  `json:"name"`	Image        string  `json:"image"`	YearStarted  int     `json:"year_started"`	FirstAlbum   string  `json:"first_album"`	Members      []string`json:"members"`	LastLocation string  `json:"last_location"`	NextLocation string  `json:"next_location"`	LastDate     string  `json:"last_date"`	NextDate     string  `json:"next_date"`
}

// ArtistsResponse represents the API response containing information about multiple artists.
type ArtistsResponse struct {
Artists []Artist `json:"artists"`
}

// Location struct represents a concert location.
type Location struct {
Name string `json:"name"`
}

// LocationsResponse represents the API response containing information about multiple concert locations.
type LocationsResponse struct {
Locations []Location `json:"locations"`
}

// Date struct represents a concert date.
type Date struct {
Date string `json:"date"`
}

// DatesResponse represents the API response containing information about multiple concert dates.
type DatesResponse struct {
Dates []Date `json:"dates"`
}

// Relation struct represents the relation between artists, dates, and locations.
type Relation struct {
ArtistID string `json:"artist_id"`
DateID string `json:"date_id"`
LocationID string `json:"location_id"`
}

// RelationsResponse represents the API response containing information about the relations.
type RelationsResponse struct {
Relations []Relation `json:"relations"`
}

```

## Interface model:

The interface models are used to define common methods that can be implemented by any concrete implementation of an entity

```bash
// Artist represents information about a band or artist.
type Artist struct {
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	YearStarted  int      `json:"year_started"`
	FirstAlbum   string   `json:"first_album"`
	Members      []string `json:"members"`
	LastLocation string   `json:"last_location"`
	NextLocation string   `json:"next_location"`
	LastDate     string   `json:"last_date"`
	NextDate     string   `json:"next_date"`
}

// Location represents a concert location.
type Location struct {
	Name string `json:"name"`
}

// Date represents a concert date.
type Date struct {
	Date string `json:"date"`
}

// Relation represents the relation between artists, dates, and locations.
type Relation struct {
	ArtistID    string `json:"artist_id"`
	DateID      string `json:"date_id"`
	LocationID  string `json:"location_id"`
}

// DataSource interface defines the methods for retrieving data.
type DataSource interface {
	GetArtists() ([]Artist, error)
	GetLocations() ([]Location, error)
	GetDates() ([]Date, error)
	GetRelations() ([]Relation, error)
}

// APIDataSource implements the DataSource interface and retrieves data from the API.
type APIDataSource struct {
	// API client configuration and other necessary fields
}

// GetArtists retrieves the list of artists from the API.
func (api *APIDataSource) GetArtists() ([]Artist, error) {
	// API implementation to fetch artists
	return nil, nil
}

// GetLocations retrieves the list of locations from the API.
func (api *APIDataSource) GetLocations() ([]Location, error) {
	// API implementation to fetch locations
	return nil, nil
}

// GetDates retrieves the list of dates from the API.
func (api *APIDataSource) GetDates() ([]Date, error) {
	// API implementation to fetch dates
	return nil, nil
}

// GetRelations retrieves the list of relations from the API.
func (api *APIDataSource) GetRelations() ([]Relation, error) {
	// API implementation to fetch relations
	return nil, nil
}

```

## Further reading and info:

Exploration on interfaces in Go, you can refer to the official [Go documentation on interfaces](https://golang.org/doc/effective_go#interfaces), ["A Tour of Go" tutorial](https://tour.golang.org/methods/9), and ["Go Interfaces in Action"](https://www.calhoun.io/interfaces-in-go/). These resources will provide us with a deeper understanding of interfaces and how to utilize them in our Go projects.
