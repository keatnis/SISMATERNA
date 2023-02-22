package routes

import (
	controller_embarazada "backendmod/controllers"
	db "backendmod/database"
	"backendmod/types"
	"backendmod/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForPuerpera(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORSS(router)
	//metodo para obtener localidades proporcionando una id del municipio
	router.HandleFunc("/RegistrarPuerpera/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := utils.StringToInt64(idAsString)
		if err != nil {
			respondWithErrorr(err, w)
			// We return, so we stop the function flow
			return
		}
		puerpera, err := controller_embarazada.ObtenerLocalidades(id)
		if err != nil {
			respondWithErrorr(err, w)
		} else {
			respondWithSuccesss(puerpera, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/RegistrarPuerpera", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncionn(w, r, func() (interface{}, error) {
			var puerpera types.Puerpera
			err := json.NewDecoder(r.Body).Decode(&puerpera)
			if err != nil {
				return nil, err
			}
			err = controller_embarazada.InsertPuerpera(puerpera)
			return true, err
		})
	}).Methods(http.MethodPost)

	router.HandleFunc("/RegistrarPuerpera", func(w http.ResponseWriter, r *http.Request) {
		municipios, err := controller_embarazada.GetMunicipios()
		if err == nil {
			respondWithSuccesss(municipios, w)
		} else {
			respondWithErrorr(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/ListaPuerpera", func(w http.ResponseWriter, r *http.Request) {
		puerpera, err := controller_embarazada.GetPuerpera()
		if err == nil {
			respondWithSuccesss(puerpera, w)
		} else {
			respondWithErrorr(err, w)
		}
	}).Methods(http.MethodGet)
}

// Habilitar el CORS
func enableCORSS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", db.AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCorss)
}
func middlewareCorss(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", db.AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

// Helper functions for respond with 200 or 500 code
func respondWithErrorr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccesss(data interface{}, w http.ResponseWriter) {
	//w, r, func() (interface{}, error
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  data,
		Error: "",
	})

}
func responderHttpConFuncionn(w http.ResponseWriter, r *http.Request, f func() (interface{}, error)) {
	datos, err := f()
	if err != nil {
		responderHttpConErrorr(err, w, r)
		return
	}
	responderHttpExitosoo(datos, w, r)
}
func responderHttpExitosoo(valor interface{}, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  valor,
		Error: "",
	})
}
func responderHttpConErrorr(err error, w http.ResponseWriter, r *http.Request) {
	log.Printf("Error al servir respuesta para %s: %v", r.RemoteAddr, err)
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  nil,
		Error: err.Error(),
	})
}

type RespuestaHttpp struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
