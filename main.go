package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/serenakm/MurderMysteryAPI/internal/app"
	"github.com/serenakm/MurderMysteryAPI/internal/routes"
)

func main() {
	// parse from command-line flags, allows user to specify parameters like port numbers with default values
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port") // 8080 as default port if no flag provided
	flag.Parse()                                               // process the command-line flag and populate the variables defined with flags, enabling the application to use the passed values

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	r := routes.SetupRoutes(app)

	server := &http.Server {
		Addr: ":8080",
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal()
	}
}