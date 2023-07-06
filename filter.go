package groupie

import (
	"fmt"
	"net/http"
	"strconv"
)

func ArtistsCreationDateFilter(artistsGroup *[]Artists, artistsData []Artists, startDate *int, endDate *int) { //Filtre l'API artists selon la date de création
	for i := 0; i < len(artistsData); i++ {
		if artistsData[i].CreationDate >= *startDate && artistsData[i].CreationDate <= *endDate {
			*artistsGroup = append(*artistsGroup, artistsData[i])
		}
	}
}

func ArtistsFirstAlbumDateFilter(artistsGroup *[]Artists, artistsData []Artists, startDate *int, endDate *int) {
	for i := 0; i < len(artistsData); i++ {
		tempArtistsData, _ := strconv.Atoi(artistsData[i].FirstAlbum[6:])
		if tempArtistsData >= *startDate && tempArtistsData <= *endDate {
			*artistsGroup = append(*artistsGroup, artistsData[i])
		}
	}
}

func ArtistsNbrMemberFilter(artistsGroup *[]Artists, artistsData []Artists, minMember *int, maxMember *int) { //Filtre l'API artists selon la date de création
	for i := 0; i < len(artistsData); i++ {
		if len(artistsData[i].Members) >= *minMember && len(artistsData[i].Members) <= *maxMember {
			*artistsGroup = append(*artistsGroup, artistsData[i])
		}
	}
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {

	filter := &Filter{
		CD: FilterParams{
			isSelected: r.FormValue("cdate"),
			From:       r.FormValue("cdate[from]"),
			To:         r.FormValue("cdate[to]"),
		},
		Members: FilterParams{
			isSelected: r.FormValue("members"),
			From:       r.FormValue("members[from]"),
			To:         r.FormValue("members[to]"),
		},
		FAlbum: FilterParams{
			isSelected: r.FormValue("falbum"),
			From:       r.FormValue("falbum[from]"),
			To:         r.FormValue("falbum[to]"),
		},
	}

	if err := filter.FilterError(); err != nil {
		fmt.Println(err)
	}

}

func (f *Filter) FilterError() error {
	return nil
}
