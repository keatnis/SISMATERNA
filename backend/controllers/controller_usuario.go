package controller_embarazada

import (
	db "backendmod/database"
	ty "backendmod/types"
)

func InsertUsuario(c ty.Usuario) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	sentenciaPreparada, err := db.Prepare("INSERT INTO usuario (clues, nombre_unidad, usuario, contraseña) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Clues, c.Nombreunidad, c.Usuario, c.Contraseña)
	if err != nil {
		return err
	}
	return err
}
func ObtenerUsuarios() ([]ty.Usuario, error) {
	//declarar un array si hay un error, retornamos si es el array esta vacío
	usuarios := []ty.Usuario{}
	bd, err := db.GetDB()

	if err != nil {
		return usuarios, err
	}
	rows, err := bd.Query(`SELECT idusuario, clues, nombre_unidad, usuario, contraseña FROM usuario;`)
	if err != nil {
		return usuarios, err
	}
	defer bd.Close()

	var usuario ty.Usuario
	//iteramos filas
	for rows.Next() {

		err := rows.Scan(&usuario.Id, &usuario.Clues, &usuario.Nombreunidad, &usuario.Usuario, &usuario.Contraseña)
		if err != nil {
			return usuarios, err
		}
		usuarios = append(usuarios, usuario)
	}
	rows.Close()
	return usuarios, nil
}
