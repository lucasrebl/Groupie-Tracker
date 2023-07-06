package groupie

type Relation struct {
	Index []RelationIndex
}

type RelationIndex struct {
	Id        int
	Locations map[string][]string `json:"datesLocations"`
}

type Dates struct {
	Index []DatesIndex
}

type DatesIndex struct {
	Id    int
	Dates []string
}

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relation     string   `json:"relations"`
}

type Locations struct {
	Index []LocationsIndex
}

type LocationsIndex struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

var API []Artists

type Filter struct {
	CD      FilterParams
	Members FilterParams
	FAlbum  FilterParams
	Result  []Artists
}
type FilterParams struct {
	isSelected string
	From       string
	To         string
}
