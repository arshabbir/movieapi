package clients

import (
	"errors"
	"fmt"
	"log"
	"math"
	"testmod/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbclient struct {
	//	m map[string]domain.Movie
	db *gorm.DB
}

type DBClient interface {
	Find(string) (*domain.Movie, error)
	UpdateByID(string, float64, string) error
	Create(m domain.Movie) error
	GetMovieByTitle(title string) (*domain.Movie, error)
	UpdateByYear(startYear int, endYear int, rating float64, generes string) error
	UpdateByRating(rating float64, opcode string, targetRating float64, generes string) error
	UpdateByGeneres(generes string, targetRating float64, targetgeneres string) error
}

func NewDBClient(username string, password string, host string, port int, dbname string) DBClient {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	d, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	log.Println("Connection to DB successful")
	d.AutoMigrate(&domain.Movie{})

	return &dbclient{db: d}
}

func (d *dbclient) UpdateByID(ID string, rating float64, generes string) error {
	log.Println("Update invoked in dbclient", ID, rating, generes)
	mov := domain.Movie{}
	res := d.db.Model(&mov).Where("movie_id = ?", ID).Updates(map[string]interface{}{"generes": generes, "rating": math.Round(rating*10) / 10})
	if res.Error != nil {
		return res.Error
	}
	log.Println("Update successful ", res.RowsAffected)
	return nil
}

func (d *dbclient) UpdateByRating(rating float64, opcode string, targetRating float64, generes string) error {
	log.Println("Update invoked in dbclient")
	mov := domain.Movie{}
	var condition string
	if opcode == "<" {
		condition = fmt.Sprintf("rating < %f", rating)
	}
	if opcode == ">" {
		condition = fmt.Sprintf("rating > %f", rating)
	}
	if opcode == "=" {
		condition = fmt.Sprintf("rating = %f", rating)
	}

	res := d.db.Model(&mov).Where(condition).Updates(map[string]interface{}{"generes": generes, "rating": math.Round(targetRating*10) / 10})
	if res.Error != nil {
		return res.Error
	}
	log.Println("Update successful ", res.RowsAffected)
	return nil
}
func (d *dbclient) UpdateByYear(startYear int, endYear int, rating float64, generes string) error {
	log.Println("Update invoked in dbclient")
	mov := domain.Movie{}
	var res *gorm.DB
	if endYear == 0 {
		res = d.db.Model(&mov).Where(" released_year =  ?", startYear).Updates(map[string]interface{}{"generes": generes, "rating": math.Round(rating*10) / 10})

	} else {
		res = d.db.Model(&mov).Where(" released_year between  ? AND ?", startYear, endYear).Updates(map[string]interface{}{"generes": generes, "rating": math.Round(rating*10) / 10})
	}
	if res.Error != nil {
		return res.Error
	}
	//
	log.Println("Update successful ", res.RowsAffected)
	return nil
}

func (d *dbclient) UpdateByGeneres(generes string, targetRating float64, targetgeneres string) error {
	log.Println("Update invoked in dbclient", targetRating, generes)
	mov := domain.Movie{}
	res := d.db.Model(&mov).Where("generes LIKE ?", fmt.Sprintf("%%%s%%", generes)).Updates(map[string]interface{}{"generes": targetgeneres, "rating": math.Round(targetRating*10) / 10})
	if res.Error != nil {
		return res.Error
	}
	log.Println("Update successful ", res.RowsAffected)

	return nil
}
func (d *dbclient) Find(title string) (*domain.Movie, error) {
	log.Println("Find invoked in DB client")
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	movie, err := d.GetMovieByTitle(title)
	if err != nil {
		log.Println("MOVIE TITLE NOT FOUND")
		return nil, errors.New("movie title not found")
	}
	return movie, nil
}

func (d *dbclient) Create(m domain.Movie) error {
	if !d.db.NewRecord(m) {
		return errors.New("error creating the new movie record in db")
	}
	d.db.Create(&m)
	return nil
}

func (d *dbclient) GetMovieByTitle(title string) (*domain.Movie, error) {
	var movie domain.Movie
	_ = d.db.Where("title = ?", title).Find(&movie)
	if movie.MovieId == "" {
		return nil, errors.New("title not found")
	}
	return &movie, nil
}
