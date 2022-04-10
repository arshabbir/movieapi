package clients

import (
	"errors"
	"fmt"
	"strconv"
	"testmod/domain"

	"github.com/eefret/gomdb"
)

type imdbClient struct {
	apiKey string
}
type ImdbClient interface {
	Get(string) (*domain.Movie, error)
}

func NewImdbClient(apiKey string) ImdbClient {
	return &imdbClient{apiKey: apiKey}
}

func (hc *imdbClient) Get(title string) (*domain.Movie, error) {
	// Perform the api call to imdb endpoint
	/*req, err := http.NewRequest("GET", hc.url, nil)
	if err != nil {
		return nil, errors.New("failed to create http request")
	}
	resp, err := hc.c.Do(req)
	if err != nil {
		return nil, errors.New("failed to perform http request to imdb endpoint")
	}

	movies := make([]domain.Movie, 0)
	if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
		return nil, errors.New("failed to decode response from imdb endpoint")
	}

	return movies, nil*/

	api := gomdb.Init(hc.apiKey)
	query := &gomdb.QueryData{Title: title, SearchType: gomdb.MovieSearch}

	res, err := api.MovieByTitle(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(res)
	// Need to convert resp into domain.Movies and send it
	//	movies := []domain.Movie{}
	m := domain.Movie{}
	m.MovieId = res.ImdbID
	m.ReleasedYear, err = strconv.Atoi(res.Year)
	if err != nil {
		return nil, errors.New("invalid releasing year")
	}
	m.Title = res.Title
	m.Rating, err = strconv.ParseFloat(res.ImdbRating, 32)
	if err != nil {
		return nil, errors.New("invalid releasing year")
	}
	m.Generes = res.Genre

	return &m, nil
}
