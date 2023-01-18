package main

import (
	"database/sql" // Interactuar con bases de datos
	"fmt"          // Imprimir mensajes y esas cosas

	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
)

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := "root"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "censos"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}
