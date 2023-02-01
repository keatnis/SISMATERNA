package controller_embarazada

import (
	db "backendmod/database"
	"backendmod/types"
)

func insertEmbarazada(c types.Embarazada) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	//_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", embarazada.NombreCompleto)
	//return err
	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO mujer (No_Expediente, Nombre,FechaNacimiento, curp, Telefono, Domicilio_Referencia, Gestas, Paras, Abortos, Cesareas, Donde_Emigro, Consulta_RiesgoPreg, Fecha_UltimoParto) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.NoExpediente, c.NombreCompleto, c.FechaNacimiento, c.Curp, c.Telefono, c.Direccion, c.Gestas, c.Paras, c.Abortos, c.Cesareas, c.DondeMigro, c.ConsultaPregestacional, c.FechaUltimoEvento)
	if err != nil {
		return err
	}
	return nil
}

func deleteVideoGame(id int64) error {

	bd, err := db.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM video_games WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func updateEmbarazada(e types.Embarazada) error {
	bd, err := db.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE video_games SET name = ?, genre = ?, year = ? WHERE id = ?",
		e.NombreCompleto, e.Direccion)
	return err
}
func GetEmbarazada() ([]types.Embarazada, error) {
	//Declare an array because if there's error, we return it empty
	embarazadas := []types.Embarazada{}
	bd, err := db.GetDB()
	if err != nil {
		return embarazadas, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query(`SELECT idMujer,Nuevo_Ingreso,Numero_Expediente, Nombre,curp,telefono,Edad, COALESCE(FechaNac, ''), domicilio_referencia, gestas, paras, abortos, cesareas, dondeEmigro, COALESCE(ConsultaPregestacional, ''), COALESCE(FechaUltimoEvento, '')  FROM mujer`)
	if err != nil {
		return embarazadas, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var embarazada types.Embarazada
		err := rows.Scan(&embarazada.Id, &embarazada.NoExpediente, &embarazada.NombreCompleto, &embarazada.Curp, &embarazada.Telefono, &embarazada.FechaNacimiento, &embarazada.Direccion, &embarazada.Gestas, &embarazada.Paras, &embarazada.Abortos, &embarazada.Cesareas, &embarazada.DondeMigro, &embarazada.ConsultaPregestacional, &embarazada.FechaUltimoEvento)
		if err != nil {
			return embarazadas, err
		}
		// and append it to the array
		embarazadas = append(embarazadas, embarazada)
	}
	return embarazadas, nil
}

func GetMunicipios() ([]types.Municipio, error) {
	municipios := []types.Municipio{}
	bd, err := db.GetDB()
	if err != nil {
		return municipios, err
	}
	rows, err := bd.Query("SELECT * from t_municipio")
	if err != nil {
		return municipios, err
	}
	//iteramos las filas
	for rows.Next() {
		var municipio types.Municipio
		err := rows.Scan(&municipio.Id_municipio, &municipio.NombreMunicipio)
		if err != nil {
			return municipios, err
		}
		municipios = append(municipios, municipio)
	}
	return municipios, nil

}

func ObtenerLocalidades(id int64) ([]types.Localidad, error) {
	localidades := []types.Localidad{}
	bd, err := db.GetDB()

	if err != nil {
		return localidades, err
	}

	defer bd.Close()
	filas, err := bd.Query(`SELECT id_localidad,localidad FROM t_localidad WHERE id_municipio=?`, id)
	if err != nil {
		return localidades, err
	}
	defer filas.Close()
	var localidad types.Localidad
	for filas.Next() {
		err := filas.Scan(&localidad.Id_localidad, &localidad.NombreLocalidad)

		if err != nil {
			return localidades, err
		}
		localidades = append(localidades, localidad)
	}
	return localidades, nil
}

func GetMunicipioById(idMunicipio string) (types.Municipio, error) {
	var idmunicipio types.Municipio
	bd, err := db.GetDB()
	if err != nil {
		return idmunicipio, err
	}

	row := bd.QueryRow("SELECT id_localidad FROM t_localidad WHERE municipio=?", idMunicipio)
	err = row.Scan(&idmunicipio.Id_municipio)
	if err != nil {
		return idmunicipio, err
	}
	// Success!
	return idmunicipio, nil
}
