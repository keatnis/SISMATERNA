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

	sentenciaPreparada, err := db.Prepare("INSERT INTO puerpera (consulta_siete, consulta_veintiocho, consulta_cuarenta, atencion_parto, lugar_parto, resolucion, no_producto, apeo, puerpera_aceptante, consulta_pregestacional) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Consultasiete, c.Consultaveinte, c.Consultacuarenta, c.Atencionparto, c.Lugarparto, c.Resolucion, c.Producto, c.Apeo, c.Aceptante, c.Pregestacional)
	if err != nil {
		return err
	}
	return err
}
func ObtenerPuerperas() ([]ty.Puerpera, error) {
	//declarar un array si hay un error, retornamos si es el array esta vacío
	puerperas := []ty.Puerpera{} 
	bd, err := db.GetDB()

	if err != nil {
		return puerperas, err
	}
	rows, err := bd.Query(`SELECT  consulta_siete,consulta_veintiocho,consulta_cuarenta,resolucion,no_producto,puerpera_aceptante,consulta_pregestacional,lugar_parto,signos_alarma,atencion_parto FROM puerpera;`)
	if err != nil {
		return puerperas, err
	}
	defer bd.Close()

	var puerpera ty.Puerpera 
	//iteramos filas
	for rows.Next() {

		err := rows.Scan(&puerpera.Consultasiete, &puerpera.Consultaveinte, &puerpera.Consultacuarenta, &puerpera.Resolucion, &puerpera.Producto, &puerpera.Aceptante, &puerpera.Pregestacional, &puerpera.Lugarparto, &puerpera.Signos, &puerpera.Atencionparto)
		if err != nil {
			return puerperas, err
		}
		puerperas = append(puerperas, puerpera)
	}
	rows.Close()
	return puerperas, nil
}
