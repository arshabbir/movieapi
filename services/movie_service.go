package services

import (
	"log"
	"testmod/clients"
	"testmod/domain"
)

type movieService struct {
	dClient    clients.DBClient
	imdbClient clients.ImdbClient
}

type MovieService interface {
	Find(string) (*domain.Movie, error)
	UpdateByID(string, float64, string) error
	UpdateByYear(startYear int, endYear int, ratings float64, generes string) error
	UpdateByRating(rating float64, opcode string, targetRating float64, generes string) error
	UpdateByGeneres(generes string, targetRating float64, targetgeneres string) error
}

func NewMovieService(dbclient clients.DBClient, imdbclient clients.ImdbClient) MovieService {
	if dbclient == nil {
		return nil
	}
	return &movieService{
		dClient:    dbclient,
		imdbClient: imdbclient,
	}
}

func (ms *movieService) Find(title string) (*domain.Movie, error) {
	resp, err := ms.dClient.Find(title)
	if err != nil && err.Error() == "movie title not found" {
		// Perform the imdb API hit
		if movie, err := ms.imdbClient.Get(title); err != nil {
			return nil, err
		} else {
			if err := ms.dClient.Create(*movie); err != nil {
				log.Println("Record  Creating error : ", movie)
			}
			return movie, nil
		}

		// store the result in localDB

	}
	log.Println("Serving from Local Database ")
	return resp, nil
}

func (ms *movieService) UpdateByID(ID string, rating float64, generes string) error {
	return ms.dClient.UpdateByID(ID, rating, generes)
}

func (ms *movieService) UpdateByYear(startYear int, endYear int, ratings float64, generes string) error {
	return ms.dClient.UpdateByYear(startYear, endYear, ratings, generes)
}

func (ms *movieService) UpdateByRating(rating float64, opcode string, targetRating float64, generes string) error {
	return ms.dClient.UpdateByRating(rating, opcode, targetRating, generes)
}

func (ms *movieService) UpdateByGeneres(generes string, targetRating float64, targetgeneres string) error {
	return ms.dClient.UpdateByGeneres(generes, targetRating, targetgeneres)
}
