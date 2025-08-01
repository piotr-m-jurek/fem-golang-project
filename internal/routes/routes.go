package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/piotr-m-jurek/fem-project/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutByID)
	return r

}
