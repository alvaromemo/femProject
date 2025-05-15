package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alvaromemo/femProject/internal/api"
	"github.com/alvaromemo/femProject/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here
	pgDb, err := store.Open()
	if err != nil {
		return nil, err
	}

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	return &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDb,
	}, nil
}

func (*Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available")
}
