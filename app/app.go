package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Bogdanov-G/authorization_app/domain"
	"github.com/Bogdanov-G/authorization_app/logger"
	"github.com/Bogdanov-G/authorization_app/service"
	"github.com/gorilla/mux"
)

// Start() starts server and make it ready to handle requests on incoming
// connections.
func Start() {

	sanityCheck()
	pool := getDBPool()

	router := mux.NewRouter()
	th := TokenHandler{
		service.NewTokenService(domain.NewTokenRepository(pool)),
	}

	router.HandleFunc("/auth/login", th.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", th.Register).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", th.Verify).Methods(http.MethodGet)

	addres := os.Getenv("SRV_ADDR")
	port := os.Getenv("SRV_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", addres, port), router)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(fmt.Sprintf("Server started as http://%s:%s", addres, port))
}

// sanityCheck() checks whether all used global environment variables have been
// presented.
func sanityCheck() {
	// All global variables required for app have to be included here.
	neededEnvs := []string{
		"SRV_ADDR",
		"SRV_PORT",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASS",
	}
	for _, env := range neededEnvs {
		if os.Getenv(env) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined!",
				env))
		}
	}
}

// getDBPool() opens database and verify that the data source name is valid via
// calling Ping.
func getDBPool() *sql.DB {
	var err error
	var pool *sql.DB

	dbAddres := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbAddres, dbPort, dbUser, dbPassword, dbName)

	pool, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = pool.Ping()
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("DB successfully connected!")

	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)

	return pool
}
