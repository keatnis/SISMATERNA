package main

import (
	db "backendmod/database"
	"backendmod/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//FOLIO DE SE
//4eUqeNLm68d
func main() {
	// Ping database
	bd, err := db.GetDB()
	if err != nil {
		log.Printf("Error with database" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error making connection to DB. Please check credentials. The error is: " + err.Error())
			return
		}
	}
	fmt.Printf("conectado correctamente ")
	// Define routes
	router := mux.NewRouter()
	routes.SetupRoutesForEmbarazada(router)
	// .. here you can define more routes
	// ...
	// for example setupRoutesForGenres(router)

	// Setup and start server
	port := ":5000"

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", port)

	log.Fatal(server.ListenAndServe())
}
