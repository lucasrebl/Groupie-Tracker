package groupie

import (
	"encoding/json"
)

func RelationJsonFormate(body *[]byte, data *Relation) {
	json.Unmarshal(*body, &data)
}

func DatesJsonFormate(body *[]byte, data *Dates) {
	json.Unmarshal(*body, &data)
}

func ArtistsJsonFormate(body *[]byte, data *[]Artists) {
	json.Unmarshal(*body, &data)
}

func LocationsJsonFormate(body *[]byte, data *Locations) {
	json.Unmarshal(*body, &data)
}
