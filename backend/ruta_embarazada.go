package main

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

/* "errors"*/
func configurarRutasEmbarazadas(enrutador *mux.Router) {
	enrutador.HandleFunc("/agregar-embarazada", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var vehiculo Embarazada
			err := json.NewDecoder(r.Body).Decode(&vehiculo)
			if err != nil {
				return nil, err
			}
			err = insertarEmbarazada(vehiculo)
			return true, err
		})
	}).Methods(http.MethodPost)
	enrutador.HandleFunc("/ListaEmbarazada", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			return obtenerVehiculos()
		})
	}).Methods(http.MethodGet)
}
