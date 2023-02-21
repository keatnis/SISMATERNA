package controller_puerpera

import (
	db "backendmod/database"
	"backendmod/types"
)

func InsertPuerpera(c types.Puerpera) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	sentenciaPreparada, err := db.Prepare("INSERT INTO puerpera (consutal_siete, consulta_veintiocho,consulta_cuarenta) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Consultasiete, c.Consultaveinte, c.Consultacuarenta)
	if err != nil {
		return err
	}
	return
}
