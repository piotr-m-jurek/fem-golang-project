package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/piotr-m-jurek/fem-project/internal/app"
	"github.com/piotr-m-jurek/fem-project/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 42069, "port of the application")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Hour,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are Live on port %d!\n", port)

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}

}
