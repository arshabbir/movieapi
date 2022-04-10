package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"testmod/domain"
	"testmod/services"
	"testmod/utils"

	"github.com/gorilla/mux"
)

type MovieController interface {
	Find(http.ResponseWriter, *http.Request)
	UpdateByRating(http.ResponseWriter, *http.Request)
	UpdateByGenere(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Ping(http.ResponseWriter, *http.Request)
}

type movieController struct {
	ms services.MovieService
}

func NewMovieController(ms services.MovieService) MovieController {
	return &movieController{ms: ms}
}

func (mc *movieController) Find(w http.ResponseWriter, r *http.Request) {
	log.Println("Invoked Find")
	vars := mux.Vars(r)
	title := vars["title"]
	resp, err := mc.ms.Find(title)

	if err != nil {
		if err := json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusInternalServerError, Message: err.Error()}); err != nil {
			log.Println("Error in encoding")
			return
		}
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Error in encoding")
		return
	}

}

func (mc *movieController) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update invoked")
	m := domain.Update{}

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusBadRequest, Message: "invalid request body " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("%#v", m)

	if m.Movieid != "" {
		log.Println("ID present Update by ID case")
		if err := mc.ms.UpdateByID(m.Movieid, m.TargetRating, m.TargetGeneres); err != nil {
			json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusBadRequest, Message: "invalid request body " + err.Error()})
			return
		}
		json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusOK, Message: "Update Successful"})
		return
	}

	if m.Ratingvector.RatingValue != 0 && m.Ratingvector.OpCode != "" {
		log.Println("Update  by Rating Value ")
		if err := mc.ms.UpdateByRating(m.Ratingvector.RatingValue, m.Ratingvector.OpCode, m.TargetRating, m.TargetGeneres); err != nil {
			json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusBadRequest, Message: "invalid request body " + err.Error()})
			return
		}
		json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusOK, Message: "Update Successful"})

		return
	}

	if m.YearVector.Startyear != 0 {
		log.Println("Update  by Year range case  ")
		if err := mc.ms.UpdateByYear(m.YearVector.Startyear, m.YearVector.Endyear, m.TargetRating, m.TargetGeneres); err != nil {
			json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusBadRequest, Message: "invalid request body " + err.Error()})
			return
		}
		json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusOK, Message: "Update Successful"})
		return
	}

	if m.GenereVector.Generes != "" {
		log.Println("Update  by Generes  case  ")
		if err := mc.ms.UpdateByGeneres(m.GenereVector.Generes, m.TargetRating, m.TargetGeneres); err != nil {
			json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusBadRequest, Message: "invalid request body " + err.Error()})
			return
		}
		json.NewEncoder(w).Encode(&utils.ApiError{Status: http.StatusOK, Message: "Update Successful"})

		return
	}

}

func (mc *movieController) UpdateByRating(w http.ResponseWriter, r *http.Request) {
	log.Println("Invoked Update")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	log.Println(vars)
}

func (mc *movieController) UpdateByGenere(w http.ResponseWriter, r *http.Request) {
	log.Println("Invoked Update")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	log.Println(vars)
}

func (mc *movieController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
