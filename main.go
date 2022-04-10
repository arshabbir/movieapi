package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"testmod/app"
	"testmod/app/controller"
	"testmod/clients"
	"testmod/services"
)

type config struct {
	dbUser     string
	dbPassword string
	dbName     string
	apiKey     string
	host       string
	dbPort     int
	appPort    int
}

func main() {
	conf := config{}

	if err := readEnv(&conf); err != nil {
		log.Println("error while reading environment variables", err)
		return
	}
	log.Println(conf)

	if err := app.NewApp(controller.NewMovieController(services.NewMovieService(clients.NewDBClient(conf.dbUser, conf.dbPassword, conf.host, conf.dbPort, conf.dbName), clients.NewImdbClient(conf.apiKey))), conf.appPort).StartApp(); err != nil {
		log.Println("Error starting the Applicaiton", err.Error())
	}
}


func readEnv(c *config) error {
	var err error
	c.dbName = os.Getenv("DBNAME")
	c.dbUser = os.Getenv("DBUSERNAME")
	c.host = os.Getenv("DBHOST")
	c.apiKey = os.Getenv("APIKEY")
	c.appPort, err = strconv.Atoi(os.Getenv("APPPORT"))
	if err != nil {
		return errors.New("application port should be a number")
	}
	c.dbPort, err = strconv.Atoi(os.Getenv("DBPORT"))
	if err != nil {
		return errors.New("dbport should be a number")
	}
	c.dbPassword = os.Getenv("DBPASSWORD")
	if c.apiKey == "" || c.dbName == "" || c.dbPassword == "" || c.host == "" || c.dbPort <= 0 || c.appPort <= 0 {
		return errors.New("error in config")
	}
	return nil
}
