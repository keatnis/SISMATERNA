package routes

import (
	controller_embarazada "backendmod/controllers"
	db "backendmod/database"
	emb "backendmod/types"

	"backendmod/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* Ejemplo de rutas
router.HandleFunc("/videogames", func(w http.ResponseWriter, r *http.Request) {
		videoGames, err := getVideoGames()
		if err == nil {
			respondWithSuccess(videoGames, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)
	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		videogame, err := getVideoGameById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(videogame, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createVideoGame(videoGame)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := updateVideoGame(videoGame)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		err = deleteVideoGame(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}

*/

func SetupRoutesForEmbarazada(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)
	//metodo para obtener localidades proporcionando una id del municipio
	router.HandleFunc("/agregar_embarazada/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := utils.StringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		embarazada, err := controller_embarazada.ObtenerLocalidades(id)

		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(embarazada, w)
			return
		}
	}).Methods(http.MethodGet)
	/* router.HandleFunc("/agregar_embarazada", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var embarazada types.Embarazada
		err := json.NewDecoder(r.Body).Decode(&embarazada)
		if err != nil {
			respondWithError(err, w)
			return
		} else {
			err := controller_embarazada.InsertEmbarazada(embarazada)
			if err != nil {
				respondWithError(err, w)
				return
			} else {
		 	respondWithSuccess(true, w)
			return
			}
		}

	}).Methods(http.MethodPost)
	*/
	router.HandleFunc("/agregar_embarazada", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var embarazada emb.Embarazada
			err := json.NewDecoder(r.Body).Decode(&embarazada)
			if err != nil {
				return nil, err
			}
			err = controller_embarazada.InsertEmbarazada(embarazada)

			return true, err
		})
	}).Methods(http.MethodPost)

	/*router.HandleFunc("/agregar_embarazada/{idMunicipio}", func(w http.ResponseWriter, r *http.Request) {
		 idAsString:= mux.BuildVarsFunc(map[string]string)
		id, err :=
		if err != nil {
			respondWithError(err, w)
			return
		}
		localidades, err := controller_embarazada.GetMunicipioById(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		} else {
			respondWithSuccess(localidades, w)
		}
	}).Methods(http.MethodGet)
	*/
	router.HandleFunc("/agregar_embarazada", func(w http.ResponseWriter, r *http.Request) {
		municipios, err := controller_embarazada.GetMunicipios()
		if err == nil {
			respondWithSuccess(municipios, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/Lista_embarazada", func(w http.ResponseWriter, r *http.Request) {
		embarazadas, err := controller_embarazada.GetEmbarazada()
		if err == nil {
			respondWithSuccess(embarazadas, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	//para PUERPERA

	router.HandleFunc("/RegistrarPuerpera", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var puerpera emb.Puerpera
			err := json.NewDecoder(r.Body).Decode(&puerpera)
			if err != nil {
				return nil, err
			}
			err = controller_embarazada.InsertPuerpera(puerpera)
			return true, err
		})
	}).Methods(http.MethodPost)

	router.HandleFunc("/ListaPuerpera", func(w http.ResponseWriter, r *http.Request) {
		embarazadas, err := controller_embarazada.ObtenerPuerperas()
		if err == nil {
			respondWithSuccess(embarazadas, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	// USUARIO
	router.HandleFunc("/Tarjetas", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var login emb.Usuario
			err := json.NewDecoder(r.Body).Decode(&login)
			if err != nil {
				return nil, err
			}
			err = controller_embarazada.InsertUsuario(login)
			return true, err
		})
	}).Methods(http.MethodPost)

	router.HandleFunc("/ListaUsuario", func(w http.ResponseWriter, r *http.Request) {
		embarazadas, err := controller_embarazada.ObtenerUsuarios()
		if err == nil {
			respondWithSuccess(embarazadas, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)
}

// Habilitar el CORS
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", db.AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func middlewareCors(next http.Handler) http.Handler {
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
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {
	//w, r, func() (interface{}, error
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  data,
		Error: "",
	})

}
func responderHttpConFuncion(w http.ResponseWriter, r *http.Request, f func() (interface{}, error)) {
	datos, err := f()
	if err != nil {
		responderHttpConError(err, w, r)
		return
	}
	responderHttpExitoso(datos, w, r)
}
func responderHttpExitoso(valor interface{}, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  valor,
		Error: "",
	})
}
func responderHttpConError(err error, w http.ResponseWriter, r *http.Request) {
	log.Printf("Error al servir respuesta para %s: %v", r.RemoteAddr, err)
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  nil,
		Error: err.Error(),
	})
}

type RespuestaHttp struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
