package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	PuertoHttp                         = ":5000"
	NombreBaseDeDatos                  = "censo.db"
	DominioPermitidoCORS               = "http://192.168.1.74:8080"
	PrefijoRutaServirContenidoEstatico = "/static/"
	DirectorioContenidoEstatico        = "../frontend/dist"
)

func main() {
	db, err := db.obtenerBaseDeDatos()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}
	// Terminar conexión al terminar función
	defer db.Close()

	// Ahora vemos si tenemos conexión
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	// Listo, aquí ya podemos usar a db!
	fmt.Printf("Conectado correctamente")

	enrutador := mux.NewRouter()
	configurarRutas(enrutador)

	servidor := &http.Server{
		Handler: enrutador,
		Addr:    PuertoHttp,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	ruta := "http://localhost" + PuertoHttp
	fmt.Printf("Escuchando en %s\n", ruta)
	log.Fatal(servidor.ListenAndServe())
}
