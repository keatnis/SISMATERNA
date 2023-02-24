package controller_embarazada

import (
	db "backendmod/database"
	ty "backendmod/types"
)

func InsertPuerpera(c ty.Puerpera) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	sentenciaPreparada, err := db.Prepare("INSERT INTO puerpera (consulta_siete, consulta_veintiocho, consulta_cuarenta, signos_alarma, atencion_parto, lugar_parto, resolucion, no_producto, apeo, puerpera_aceptante, consulta_pregestacional) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Consultasiete, c.Consultaveinte, c.Consultacuarenta, c.Signos, c.Atencionparto, c.Lugarparto, c.Resolucion, c.Producto, c.Apeo, c.Aceptante, c.Pregestacional)
	if err != nil {
		return err
	}
	return err
}
