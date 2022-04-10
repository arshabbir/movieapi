package domain

import "github.com/jinzhu/gorm"

type Movie struct {
	gorm.Model
	Title        string  `json:"title"`
	ReleasedYear int     `json:"releasedyear"`
	Rating       float64 `json:"rating"`
	MovieId      string  `json:"id"`
	Generes      string  `json:"generes"`
}

type Update struct {
	Movieid       string           `json:"id"`
	YearVector    YearCondition    `json:"yearvector"`
	Ratingvector  RatingConditon   `json:"ratingvector"`
	GenereVector  GeneresCondition `json:"generesvector"`
	TargetGeneres string           `json:"targetgeneres"`
	TargetRating  float64          `json:"targetratings"`
}

type YearCondition struct {
	Startyear int `json:"startyear"`
	Endyear   int `json:"endyear"`
}
type RatingConditon struct {
	RatingValue float64 `json:"rating"`
	OpCode      string  `json:"opcode"`
}

type GeneresCondition struct {
	Generes string `json:"generes"`
}
