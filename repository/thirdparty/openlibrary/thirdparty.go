package openlibrary

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/openlibrary"
)

type OpenLibraryAPI struct {
	httpClient http.Client
}

func NewAPI() openlibrary.Repository {
	return &OpenLibraryAPI{
		httpClient: http.Client{},
	}
}

func (openLibraryAPI *OpenLibraryAPI) Fetch(links []string) ([]openlibrary.Domain, error) {
	var results []openlibrary.Domain

	for _, link := range links {
		response, _ := http.Get(link)
		responseData, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()

		var data Response

		_ = json.Unmarshal(responseData, &data)

		publicationYear, _ := strconv.Atoi(data.PublicationDate)

		var authors []string
		for _, dataAuthor := range data.Authors {
			authors = append(authors, dataAuthor["key"])
		}

		// fmt.Println(responseToDomain(data, authors, publicationYear))

		results = append(results, responseToDomain(data, authors, publicationYear))
	}

	return results, nil
}
