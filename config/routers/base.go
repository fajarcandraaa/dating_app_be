package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dbConfig "github.com/fajarcandraaa/dating_app_be/config/database"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Serve struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Serve) Initialize(dbDriver, dbUser, dbPass, dbPort, dbHost, dbName string) {
	var err error

	registry := dbConfig.SetMigrationTable()

	if dbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		s.DB, err = gorm.Open(mysql.Open(DBURL))
		if err != nil {
			fmt.Printf("Cannot connect to %s database", dbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", dbDriver)
		}
	}
	if dbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPass)
		s.DB, err = gorm.Open(postgres.Open(DBURL))
		if err != nil {
			fmt.Printf("Cannot connect to %s database", dbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", dbDriver)
		}
	}

	//Migration proccess
	s.DB.Debug().AutoMigrate(registry...) //database migration

	s.Router = mux.NewRouter()

	s.initializeRoutes()
}

// Run is used to execute connection and run our service
func (s *Serve) Run() {
	port := os.Getenv("APP_PORT")
	fmt.Println("Listening to port ", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Router))
}
