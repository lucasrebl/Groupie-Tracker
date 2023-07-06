package main

import (
	"fmt"
	"groupie"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	//GET API BODY
	relationResp, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	relationBody, _ := ioutil.ReadAll(relationResp.Body)

	datesResp, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	datesBody, _ := ioutil.ReadAll(datesResp.Body)

	artistsResp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	artistsBody, _ := ioutil.ReadAll(artistsResp.Body)

	locationsResp, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	locationsBody, _ := ioutil.ReadAll(locationsResp.Body)
	//

	//API-BODY to struc and print all
	var relationData groupie.Relation
	groupie.RelationJsonFormate(&relationBody, &relationData)

	var datesData groupie.Dates
	groupie.DatesJsonFormate(&datesBody, &datesData)

	var artistsData []groupie.Artists
	groupie.ArtistsJsonFormate(&artistsBody, &artistsData)

	var locationsData groupie.Locations
	groupie.LocationsJsonFormate(&locationsBody, &locationsData)

	fmt.Println("relation : ", "\n", relationData)                                                                                                                 //Imprime toute la structure Relation
	fmt.Println("=============", "\n", "dates : ", "\n", datesData)                                                                                                //Imprime toute la structure Dates
	fmt.Println("=============", "\n", "artists : ", "\n", artistsData)                                                                                            //Imprime toute la structure Artists
	fmt.Println("=============", "\n", "locations : ", "\n", locationsData)                                                                                        //Imprime toute la structure Location
	fmt.Println("=============", "\n", "locations (uniquement l'index 0): ", "\n", "Id : ", locationsData.Index[0].Id, " | ", locationsData.Index[0].Locations[0]) //Imprime une partie de la structure location
	//

	//FILTRAGE API ARTISTS DATE CREATION
	var artistsGroupCreationDate []groupie.Artists //résultat du filtre

	artistsGroupCreationStartDate := 1990 //date de création
	artistsGroupCreationEndDate := 1995

	groupie.ArtistsCreationDateFilter(&artistsGroupCreationDate, artistsData, &artistsGroupCreationStartDate, &artistsGroupCreationEndDate) //appel de la fonction du filtre

	fmt.Println("Date debut : ", artistsGroupCreationStartDate, "date fin : ", artistsGroupCreationEndDate, "| structure trier : ", artistsGroupCreationDate) //print la struct filtré
	//

	//FILTRAGE API ARTISTS DATE PREMIER ALBUMS
	var artistsGroupFirstAlbumDate []groupie.Artists //résultat du filtre

	artistsGroupFirstAlbumStartDate := 1970 //date de création
	artistsGroupFirstAlbumEndDate := 1980

	groupie.ArtistsFirstAlbumDateFilter(&artistsGroupFirstAlbumDate, artistsData, &artistsGroupFirstAlbumStartDate, &artistsGroupFirstAlbumEndDate) //appel de la fonction du filtre

	fmt.Println("Date debut : ", artistsGroupFirstAlbumStartDate, "date fin : ", artistsGroupFirstAlbumEndDate, "| structure trier : ", artistsGroupFirstAlbumDate) //print la struct filtré
	//

	//FILTRAGE API MEMBER
	var artistsGroupNbrMember []groupie.Artists //résultat du filtre

	artistsGroupNbrMinMember := 2 //nombre de membre
	artistsGroupNbrMaxMember := 3

	groupie.ArtistsNbrMemberFilter(&artistsGroupNbrMember, artistsData, &artistsGroupNbrMinMember, &artistsGroupNbrMaxMember) //appel de la fonction du filtre

	fmt.Println("| structure trier : ", artistsGroupNbrMember) //print la struct filtré
	//

	//TEST ZONE
	//

	//FRONT END
	printedArtists := artistsData
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		groupie.FilterHandler(w, r)
		template := template.Must(template.ParseFiles("page/accueil.html", "Static/templates/header.html", "Static/templates/footer.html"))

		DateCheckbox := r.Form["cdate"]
		fmt.Print(DateCheckbox)

		StrDateFrom := r.Form["cdate[from]"]
		fmt.Print(StrDateFrom)

		StrDateTo := r.Form["cdate[to]"]
		fmt.Print(StrDateTo)

		AlbumCheckbox := r.Form["falbum"]
		fmt.Print(AlbumCheckbox)

		StrAlbumFrom := r.Form["falbum[from]"]
		fmt.Print(StrAlbumFrom)

		StrAlbumTo := r.Form["falbum[to]"]
		fmt.Print(StrAlbumTo)

		MembersCheckbox := r.Form["members"]
		fmt.Print(MembersCheckbox)

		StrMembersFrom := r.Form["members[from]"]
		fmt.Print(StrMembersFrom)

		StrMembersTo := r.Form["members[to]"]
		fmt.Print(StrMembersTo)

		Locations := r.Form["locations"]
		fmt.Print(Locations)

		Location := r.Form["location"]
		fmt.Print(strings.Join(Location[:], " "))

		fmt.Println(len(AlbumCheckbox), len(StrAlbumFrom), len(StrAlbumTo))

		if (len(DateCheckbox) != 0) && (len(StrDateFrom) != 0) && (len(StrDateTo) != 0) {
			if (StrDateFrom[0] != "") && (StrDateTo[0] != "") {
				fmt.Println("trie date creation")
				artistsGroupCreationStartDateTemp, _ := strconv.Atoi(StrDateFrom[0])
				artistsGroupCreationStartDate = artistsGroupCreationStartDateTemp

				artistsGroupCreationEndDateTemp, _ := strconv.Atoi(StrDateTo[0])
				artistsGroupCreationEndDate = artistsGroupCreationEndDateTemp

				var artistsGroupCreationDateTemp []groupie.Artists

				groupie.ArtistsCreationDateFilter(&artistsGroupCreationDateTemp, artistsData, &artistsGroupCreationStartDate, &artistsGroupCreationEndDate)
				printedArtists = artistsGroupCreationDateTemp
				fmt.Println(artistsGroupCreationEndDate, artistsGroupCreationStartDate)
			}
		} else if (len(AlbumCheckbox) != 0) && (len(StrAlbumFrom) != 0) && (len(StrAlbumTo) != 0) {
			if (StrAlbumFrom[0] != "") && (StrAlbumTo[0] != "") {
				fmt.Println("trie date premierAlbum")
				artistsGroupFirstAlbumStartDateTemp, _ := strconv.Atoi(StrAlbumFrom[0])
				artistsGroupFirstAlbumStartDate = artistsGroupFirstAlbumStartDateTemp

				artistsGroupFirstAlbumEndDateTemp, _ := strconv.Atoi(StrAlbumTo[0])
				artistsGroupFirstAlbumEndDate = artistsGroupFirstAlbumEndDateTemp

				var artistsGroupFirstAlbumDateTemp []groupie.Artists

				groupie.ArtistsFirstAlbumDateFilter(&artistsGroupFirstAlbumDateTemp, artistsData, &artistsGroupFirstAlbumStartDate, &artistsGroupFirstAlbumEndDate) //appel de la fonction du filtre

				printedArtists = artistsGroupFirstAlbumDateTemp
				fmt.Println(artistsGroupFirstAlbumStartDate, artistsGroupFirstAlbumEndDate)
			}
		} else if (len(MembersCheckbox) != 0) && (len(StrMembersFrom) != 0) && (len(StrMembersTo) != 0) {
			if (StrMembersFrom[0] != "") && (StrMembersTo[0] != "") {
				fmt.Println("trie membre")
				artistsGroupNbrMinMemberTemp, _ := strconv.Atoi(StrMembersFrom[0])
				artistsGroupNbrMinMember = artistsGroupNbrMinMemberTemp

				artistsGroupNbrMaxMemberTemp, _ := strconv.Atoi(StrMembersTo[0])
				artistsGroupNbrMaxMember = artistsGroupNbrMaxMemberTemp

				var artistsGroupMemberTemp []groupie.Artists

				groupie.ArtistsNbrMemberFilter(&artistsGroupMemberTemp, artistsData, &artistsGroupNbrMinMember, &artistsGroupNbrMaxMember) //appel de la fonction du filtre

				printedArtists = artistsGroupMemberTemp
				fmt.Println(artistsGroupNbrMinMember, artistsGroupNbrMaxMember)
			}

		} else if (len(DateCheckbox) == 0) && (len(AlbumCheckbox) == 0) && (len(StrDateFrom) != 0) && (len(StrDateTo) != 0) && (len(StrAlbumFrom) != 0) && (len(StrAlbumTo) != 0) {
			printedArtists = artistsData
			fmt.Println("non")
		}

		if r.Method == http.MethodGet {
			template.Execute(w, printedArtists)
			return
		} else if r.Method == http.MethodPost {
			r.ParseForm()

			re := regexp.MustCompile("(?i)" + r.FormValue("query"))
			fmt.Println(re)

			filteredData := []groupie.Artists{}

			for _, d := range printedArtists {

				//filtre par nom et par membres
				if re.MatchString(d.Name) || re.MatchString(strings.Join(d.Members[:], " ")) {
					filteredData = append(filteredData, d)
				}
			}
			// searchItems.Query
			template.Execute(w, filteredData)
			// fmt.Println(filteredData)

		}
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("page/artiste.html", "Static/templates/header.html", "Static/templates/footer.html"))
		if r.Method != http.MethodPost {
			template.Execute(w, artistsData)
			return
		}
	})
	//

	//Server launch
	fs := http.FileServer(http.Dir("Static/"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fs))
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	//
}
