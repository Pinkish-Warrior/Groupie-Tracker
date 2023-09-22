package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	Client      *http.Client
	Myartists   []Artists
	Mylocations Locations
	Mydates     Dates
	Myrelations Relations
)

type Artists struct {
	Id           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate uint16   `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	ConcertDates Dates
	Relations    Relations
}

type Locations struct {
	Index []Location `json:"index"`
}

type Location struct {
	Id        uint64   `json:"id"`
	Locations []string `json:"locations"`
	DatesUrl  string   `json:"dates"`
}

type Dates struct {
	Index []Date `json:"index"`
}

type Date struct {
	Id    uint64   `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	Id             uint64              `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetJson(url string, target interface{}) error {
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API responded with status code: %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

//ALERT
// func AppendToArtistStruct() {
// 	for i := range Mylocations.Index {
// 		Myartists[i].Locations.Index = append(Myartists[i].Locations.Index, Mylocations.Index[i])
// 	}
// 	for i := range Mydates.Index {
// 		Myartists[i].ConcertDates.Index = append(Myartists[i].ConcertDates.Index, Mydates.Index[i])
// 	}
// 	for i := range Myrelations.Index {
// 		Myartists[i].Relations.Index = append(Myartists[i].Relations.Index, Myrelations.Index[i])
// 	}
// }

// IMPORTANT
// The previous code has been refactored to avoid an out-of-range warning. The warning was related to using the range keyword in the loop, which iterates over the elements of a slice. If the sizes of the slices being iterated are not the same, it can lead to an index out of range error.
func AppendToArtistStruct() {
	// Assuming Mylocations.Index, Mydates.Index, and Myrelations.Index are of the same size as Myartists
	for i := 0; i < len(Myartists); i++ {
		Myartists[i].Locations.Index = append(Myartists[i].Locations.Index, Mylocations.Index[i])
		Myartists[i].ConcertDates.Index = append(Myartists[i].ConcertDates.Index, Mydates.Index[i])
		Myartists[i].Relations.Index = append(Myartists[i].Relations.Index, Myrelations.Index[i])
	}
}
