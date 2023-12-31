package main

import (
	"groupie/tools"
	"testing"
)

func TestGetArtist(t *testing.T) {
	art := tools.GetArtist()
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test case 1",
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(art)
			if got != tt.want {
				t.Errorf("GetArtist() a retourné %d éléments, attendu %d éléments", got, tt.want)
			}
		})
	}
}

func TestGetArtistLocation(t *testing.T) {
	actual := tools.GetLocation() // Remplacez par l'appel réel à la fonction GetLocation()
	expected := [][]string{{"north_carolina-usa", "georgia-usa", "los_angeles-usa", "saitama-japan", "osaka-japan", "nagoya-japan", "penrose-new_zealand", "dunedin-new_zealand"},
		{"playa_del_carmen-mexico", "papeete-french_polynesia", "noumea-new_caledonia"},
		{"london-uk", "lausanne-switzerland", "lyon-france"},
		{"victoria-australia", "new_south_wales-australia", "queensland-australia", "auckland-new_zealand", "yogyakarta-indonesia", "bratislava-slovakia", "budapest-hungary", "minsk-belarus"},
		{"california-usa", "nevada-usa", "georgia-usa"},
		{"california-usa", "sao_paulo-brazil", "san_isidro-argentina"},
		{"california-usa", "arizona-usa texas-usa"},
		{"stockholm-sweden", "werchter-belgium", "lisbon-portugal", "bilbao-spain", "georgia-usa", "bogota-colombia", "sao_paulo-brazil"},
		{"new_york-usa", "dusseldorf-germany", "aarhus-denmark", "manchester-uk"},
		{"frankfurt-germany", "berlin-germany", "stockholm-sweden", "copenhagen-denmark", "werchter-belgium"},
		{"doha-qatar", "minnesota-usa", "illinois-usa", "california-usa", "mumbai-india"},
		{"abu_dhabi-united_arab_emirates", "new_york-usa", "pennsylvania-usa"},
		{"westcliff_on_sea-uk", "merkers-germany", "illinois-usa"},
		{"nevada-usa", "arizona-usa", "california-usa"},
		{"london-uk new_york-usa", "maine-usa"}, {"berlin-germany", "copenhagen-denmark", "aarhus-denmark", "gothenburg-sweden"},
		{"florida-usa", "south_carolina-usa", "north_carolina-usa"},
		{"pagney_derriere_barine-france", "hamburg-germany", "boulogne_billancourt-france", "nimes-france", "sion-switzerland", "ostrava-czechia", "klagenfurt-austria", "freyming_merlebach-france"},
		{"nevada-usa", "london-uk", "manchester-uk"},
		{"zaragoza-spain", "madrid-spain barcelona-spain"},
		{"rio_de_janeiro-brazil", "recife-brazil"},
		{"california-usa", "leipzig-germany", "salem-germany", "monchengladbach-germany", "cuxhaven-germany", "skanderborg-denmark", "amsterdam-netherlands", "burriana-spain", "oulu-finland", "budapest-hungary", "napoca-romania"},
		{"riyadh-saudi_arabia", "rio_de_janeiro-brazil", "canton-usa", "quebec-canada", "new_york-usa", "california-usa", "las_vegas-usa", "mexico_city-mexico", "monterrey-mexico", "del_mar-usa", "berlin-germany", "lisbon-portugal"},
		{"washington-usa", "west_melbourne-australia", "amsterdam-netherlands", "paris-france", "manchester-uk", "missouri-usa", "chicago-usa", "birmingham-uk", "copenhagen-denmark"},
		{"west_melbourne-australia", "sydney-australia", "madison-usa", "toronto-canada", "cleveland-usa new_york-usa", "boston-usa", "texas-usa utah-usa", "california-usa"},
		{"manchester-uk", "glasgow-uk", "dublin-ireland", "cardiff-uk", "london-uk", "aberdeen-uk", "stockholm-sweden", "madrid-spain", "paris-france", "warsaw-poland", "berlin-germany", "milan-italy"},
		{"michigan-usa", "new_hampshire-usa", "new_york-usa", "warsaw-poland", "sochaux-france", "lyon-france", "eindhoven-netherlands", "oslo-norway", "amsterdam-netherlands", "colorado-usa"},
		{"jakarta-indonesia", "huizhou-china", "texas-usa", "michigan-usa", "changzhou-china", "hong_kong-china", "colorado-usa", "sanya-china", "aalborg-denmark", "washington-usa", "new_york-usa", "toronto-canada", "seattle-usa"},
		{"omaha-usa", "kansas_city-usa", "st_louis-usa", "indianapolis-usa", "rosemont-usa", "grand_rapids-usa", "toronto-usa", "montreal-usa", "newark-usa", "uniondale-usa", "philadelphia-usa", "hershey-usa", "pittsburgh-usa", "washington-usa", "columbia-usa"},
		{"santiago-chile", "sao_paulo-brazil", "los_angeles-usa", "houston-usa", "atlanta-usa", "new_orleans-usa", "philadelphia-usa", "london-uk", "frauenfeld-switzerland", "turku-finland"},
		{"las_vegas-usa", "brooklyn-usa", "boston-usa", "washington-usa", "philadelphia-usa", "montreal-canada", "toronto-usa", "new_york-usa"},
		{"frankfurt-germany", "berlin-germany", "stockholm-sweden", "copenhagen-denmark", "imola-italy", "vienna-austria", "london-uk", "krakow-poland", "budapest-hungary", "zurich-switzerland"},
		{"philadelphia-usa", "paris-france", "amityville-usa", "chicago-usa", "minneapolis-usa", "detroit-usa"},
		{"oakland-usa", "charlotte-usa", "los_angeles-usa", "berlin-germany", "houston-usa", "chicago-usa", "inglewood-usa", "madrid-spain"},
		{"windsor-canada", "brooklyn-usa", "birmingham-uk", "cincinnati-usa", "anaheim-usa", "chicago-usa"},
		{"manila-philippines", "mumbai-india", "auckland-new_zealand", "brisbane-australia", "melbourne-australia", "sydney-australia"},
		{"bogota-colombia", "rio_de_janeiro-brazil", "sao_paulo-brazil", "santiago-chile", "san_isidro-argentina", "lima-peru"},
		{"paris-france", "groningen-netherlands", "antwerp-belgium", "vienna-austria", "glasgow-uk", "london-uk"},
		{"mexico_city-mexico", "pico_rivera-usa", "chicago-usa", "boston-usa", "philadelphia-usa"},
		{"london-uk", "berwyn-usa", "georgia-usa", "new_york-usa", "dallas-usa", "houston-usa"},
		{"california-usa", "birmingham-uk", "brixton-uk", "london-uk", "rotselaar-belgium"},
		{"california-usa", "rio_de_janeiro-brazil", "los_angeles-usa", "alabama-usa", "massachusetts-usa", "athens-greece", "florence-italy", "landgraaf-netherlands"},
		{"sydney-australia", "melbourne-australia", "burswood-australia", "wellington-new_zealand", "abu_dhabi-united_arab_emirates"},
		{"madrid-spain", "seville-spain", "los_angeles-usa", "bangkok-thailand", "manila-philippines", "taipei-taiwan", "hong_kong-china", "seoul-south_korea"},
		{"munich-germany", "mannheim-germany", "san_francisco-usa", "santiago-chile", "buenos_aires-argentina", "porto_alegre-brazil", "sao_paulo-brazil", "belo_horizonte-brazil"},
		{"california-usa", "sao_paulo-brazil", "porto_alegre-brazil", "la_plata-argentina", "london-uk"},
		{"dubai-united_arab_emirates", "willemstad-netherlands_antilles", "florida-usa", "florida-usa", "mexico_city-mexico", "santiago-chile", "sao_paulo-brazil", "brasilia-brazil"},
		{"texas-usa", "oklahoma-usa", "california-usa", "illinois-usa", "scheessel-germany", "st_gallen-switzerland", "gdynia-poland", "arras-france"},
		{"washington-usa", "california-usa", "california-usa", "arizona-usa", "florida-usa"},
		{"mexico_city-mexico", "sao_paulo-brazil", "buenos_aires-argentina", "santiago-chile", "lima-peru"},
		{"bogota-colombia", "san_jose-costa_rica", "nevada-usa", "massachusetts-usa", "massachusetts-usa", "nickelsdorf-austria", "milan-italy", "lisbon-portugal"},
		{"oregon-usa", "vancouver-canada", "nevada-usa", "colorado-usa", "munich-germany", "prague-czechia", "milan-italy"}}
	if len(actual) != len(expected) {
		t.Errorf("Le nombre de localisations réelles (%d) ne correspond pas au nombre attendu (%d)", len(actual), len(expected))
	}
}
