package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type University struct {
	Domains       []string `json:"domains"`
	StateProvince *string  `json:"state-province"`
	Country       string   `json:"country"`
	AlphaTwoCode  string   `json:"alpha_two_code"`
	WebPages      []string `json:"web_pages"`
	Name          string   `json:"name"`
}

type Country struct {
	Name         string       `json:"name"`
	Universities []University `json:"universities"`
}

// Fetch all universities in all countries concurrently using GoParallel from hipolabs public universities api
func main() {
	start := time.Now()

	// Example of sequential execution, 23.09 seconds
	for _, country := range allCountries {
		getUniversities(country)
	}
	elapsed := time.Since(start)
	fmt.Println("Sequential Run Time elapsed:", elapsed)

	// Use GoParallel to fetch the universities for all countries concurrently

	// 3 Levels of Parallelism, 8.08 seconds
	start = time.Now()
	GoParallel.ProcessSlice(allCountries, getUniversities, 3)
	elapsed = time.Since(start)
	fmt.Println("3 Levels of Parallelism run, time took:", elapsed)

	// 6 Levels of Parallelism, 4.42 seconds
	start = time.Now()
	GoParallel.ProcessSlice(allCountries, getUniversities, 6)
	elapsed = time.Since(start)
	fmt.Println("6 Levels of Parallelism run, time took:", elapsed)

}

func getUniversities(country *Country) {
	// Send an HTTP GET request to the API
	resp, err := http.Get(fmt.Sprintf("http://universities.hipolabs.com/search?country=%s", url.QueryEscape(country.Name)))
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Unmarshal the JSON response into a slice of University structs
	var universities []University
	err = json.Unmarshal(body, &universities)
	if err != nil {
		fmt.Println("Error unmarshaling the JSON response:", err)
		return
	}
	// Add universities to the country struct
	country.Universities = universities
}

var allCountries = []*Country{
	{Name: "Austria"},
	{Name: "Guinea"},
	{Name: "Israel"},
	{Name: "Korea, Republic of"},
	{Name: "Colombia"},
	{Name: "Guatemala"},
	{Name: "Somalia"},
	{Name: "Dominica"},
	{Name: "El Salvador"},
	{Name: "Benin"},
	{Name: "Italy"},
	{Name: "Saint Vincent and the Grenadines"},
	{Name: "Moldova, Republic of"},
	{Name: "Nepal"},
	{Name: "Iraq"},
	{Name: "Egypt"},
	{Name: "Luxembourg"},
	{Name: "Macao"},
	{Name: "Uzbekistan"},
	{Name: "Kenya"},
	{Name: "Burkina Faso"},
	{Name: "Montserrat"},
	{Name: "Japan"},
	{Name: "Belize"},
	{Name: "Sweden"},
	{Name: "Singapore"},
	{Name: "Saint Kitts and Nevis"},
	{Name: "Yemen"},
	{Name: "Libya"},
	{Name: "Belarus"},
	{Name: "Taiwan"},
	{Name: "Thailand"},
	{Name: "Mauritius"},
	{Name: "Cape Verde"},
	{Name: "Barbados"},
	{Name: "Brunei Darussalam"},
	{Name: "Cuba"},
	{Name: "Suriname"},
	{Name: "Iran"},
	{Name: "Maldives"},
	{Name: "Namibia"},
	{Name: "Philippines"},
	{Name: "Jordan"},
	{Name: "Syrian Arab Republic"},
	{Name: "Guadeloupe"},
	{Name: "Chad"},
	{Name: "Rwanda"},
	{Name: "Korea, Democratic People's Republic of"},
	{Name: "Belgium"},
	{Name: "Kyrgyzstan"},
	{Name: "Nicaragua"},
	{Name: "New Caledonia"},
	{Name: "Tunisia"},
	{Name: "Congo, the Democratic Republic of the"},
	{Name: "Afghanistan"},
	{Name: "Ghana"},
	{Name: "Faroe Islands"},
	{Name: "Turks and Caicos Islands"},
	{Name: "Paraguay"},
	{Name: "Dominican Republic"},
	{Name: "Papua New Guinea"},
	{Name: "Poland"},
	{Name: "Guam"},
	{Name: "Kuwait"},
	{Name: "Morocco"},
	{Name: "Norway"},
	{Name: "Palestine, State of"},
	{Name: "Serbia"},
	{Name: "Viet Nam"},
	{Name: "Oman"},
	{Name: "Switzerland"},
	{Name: "China"},
	{Name: "Peru"},
	{Name: "Grenada"},
	{Name: "Georgia"},
	{Name: "Holy See (Vatican City State)"},
	{Name: "Ireland"},
	{Name: "Kazakhstan"},
	{Name: "Andorra"},
	{Name: "Saudi Arabia"},
	{Name: "Lesotho"},
	{Name: "Djibouti"},
	{Name: "Uganda"},
	{Name: "Estonia"},
	{Name: "Senegal"},
	{Name: "Portugal"},
	{Name: "Germany"},
	{Name: "Bulgaria"},
	{Name: "BR"},
	{Name: "Latvia"},
	{Name: "Brazil"},
	{Name: "Botswana"},
	{Name: "Seychelles"},
	{Name: "Lao People's Democratic Republic"},
	{Name: "UK"},
	{Name: "Greece"},
	{Name: "Tanzania, United Republic of"},
	{Name: "Madagascar"},
	{Name: "Burundi"},
	{Name: "Bermuda"},
	{Name: "India"},
	{Name: "Nigeria"},
	{Name: "Algeria"},
	{Name: "Trinidad and Tobago"},
	{Name: "Panama"},
	{Name: "Uruguay"},
	{Name: "South Sudan"},
	{Name: "Mozambique"},
	{Name: "Canada"},
	{Name: "Eritrea"},
	{Name: "Malta"},
	{Name: "Gambia"},
	{Name: "Sudan"},
	{Name: "Qatar"},
	{Name: "Réunion"},
	{Name: "Hungary"},
	{Name: "Samoa"},
	{Name: "Slovakia"},
	{Name: "French Guiana"},
	{Name: "Niue"},
	{Name: "Ethiopia"},
	{Name: "Bahamas"},
	{Name: "Liberia"},
	{Name: "Mexico"},
	{Name: "Romania"},
	{Name: "Mauritania"},
	{Name: "Montenegro"},
	{Name: "Puerto Rico"},
	{Name: "Tajikistan"},
	{Name: "Turkmenistan"},
	{Name: "Guyana"},
	{Name: "Chile"},
	{Name: "Antigua and Barbuda"},
	{Name: "Indonesia"},
	{Name: "Lebanon"},
	{Name: "Mongolia"},
	{Name: "Togo"},
	{Name: "Jamaica"},
	{Name: "Spain"},
	{Name: "Hong Kong"},
	{Name: "Sierra Leone"},
	{Name: "Costa Rica"},
	{Name: "Pakistan"},
	{Name: "Niger"},
	{Name: "Congo"},
	{Name: "Myanmar"},
	{Name: "Malaysia"},
	{Name: "Bahrain"},
	{Name: "Honduras"},
	{Name: "Croatia"},
	{Name: "US"},
	{Name: "Lithuania"},
	{Name: "Bosnia and Herzegovina"},
	{Name: "Fiji"},
	{Name: "Zimbabwe"},
	{Name: "New Zealand"},
	{Name: "Cyprus"},
	{Name: "Côte d'Ivoire"},
	{Name: "Denmark"},
	{Name: "Bangladesh"},
	{Name: "Cayman Islands"},
	{Name: "Bhutan"},
	{Name: "Central African Republic"},
	{Name: "Mali"},
	{Name: "Netherlands"},
	{Name: "Ukraine"},
	{Name: "Venezuela, Bolivarian Republic of"},
	{Name: "Azerbaijan"},
	{Name: "Russian Federation"},
	{Name: "Turkey"},
	{Name: "Ecuador"},
	{Name: "Equatorial Guinea"},
	{Name: "United States"},
	{Name: "Monaco"},
	{Name: "Czech Republic"},
	{Name: "Liechtenstein"},
	{Name: "France"},
	{Name: "Armenia"},
	{Name: "Finland"},
	{Name: "Saint Lucia"},
	{Name: "Argentina"},
	{Name: "United Kingdom"},
	{Name: "Sri Lanka"},
	{Name: "Gabon"},
	{Name: "Haiti"},
	{Name: "Angola"},
	{Name: "United Arab Emirates"},
	{Name: "Vietnam"},
	{Name: "Australia"},
	{Name: "Cambodia"},
	{Name: "Iceland"},
	{Name: "Albania"},
	{Name: "Greenland"},
	{Name: "Bolivia, Plurinational State of"},
	{Name: "Swaziland"},
	{Name: "Zambia"},
	{Name: "French Polynesia"},
	{Name: "Kosovo"},
	{Name: "Slovenia"},
	{Name: "Virgin Islands, British"},
	{Name: "Cameroon"},
	{Name: "Malawi"},
	{Name: "North Macedonia"},
	{Name: "South Africa"},
	{Name: "San Marino"},
}
