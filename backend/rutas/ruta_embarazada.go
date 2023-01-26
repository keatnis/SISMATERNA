package rutas

import (
	"encoding/json"
	"net/http"
	"sismat/controlador_embarazada"
	"sismat/utiles"

	"github.com/gorilla/mux"
)

/* "errors"*/
func ConfigurarRutasEmbarazadas(enrutador *mux.Router) {
	enrutador.HandleFunc("/agregar-embarazada", func(w http.ResponseWriter, r *http.Request) {
		utiles.ResponderHttpConFuncion(w, r, func() (interface{}, error) {
			var vehiculo controlador_embarazada.Embarazada
			err := json.NewDecoder(r.Body).Decode(&vehiculo)
			if err != nil {
				return nil, err
			}
			err = controlador_embarazada.InsertarEmbarazada(vehiculo)
			return true, err
		})
	}).Methods(http.MethodPost)
	enrutador.HandleFunc("/ListaEmbarazada", func(w http.ResponseWriter, r *http.Request) {
		utiles.ResponderHttpConFuncion(w, r, func() (interface{}, error) {
			return controlador_embarazada.ObtenerVehiculos()
		})
	}).Methods(http.MethodGet)
}
