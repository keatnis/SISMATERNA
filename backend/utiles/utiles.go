package utiles

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type RespuestaHttp struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func ObtenerVariableEnteraDeUrl(variablesDePeticion map[string]string, variable string) (int64, error) {
	numero, err := strconv.ParseInt(variablesDePeticion[variable], 0, 64)
	if err != nil {
		return 0, err
	}
	return numero, err
}

func ResponderHttpExitoso(valor interface{}, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  valor,
		Error: "",
	})
}
func ResponderHttpConError(err error, w http.ResponseWriter, r *http.Request) {
	log.Printf("Error al servir respuesta para %s: %v", r.RemoteAddr, err)
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  nil,
		Error: err.Error(),
	})
}

func Responder(datos interface{}, err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ResponderHttpConError(err, w, r)
		return
	}
	ResponderHttpExitoso(datos, w, r)
}
func ResponderHttpConFuncion(w http.ResponseWriter, r *http.Request, f func() (interface{}, error)) {
	datos, err := f()
	if err != nil {
		ResponderHttpConError(err, w, r)
		return
	}
	ResponderHttpExitoso(datos, w, r)
}
