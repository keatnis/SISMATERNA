package controller_embarazada

import (
	db "backendmod/database"
	"backendmod/types"
)

func InsertEmbarazada(c types.Embarazada) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	//_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", embarazada.NombreCompleto)
	//return err
	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO mujer (noExpediente, Nombre,FechaNacimiento, curp, Telefono, Domicilio_Referencia, Gestas, Paras, Aborto, Cesareas, emigro, Consulta_RiesgoPreg, Fecha_UltimoParto,Fecha_Ult_Mestruacion,PFP,Fecha_Consulta,Fecha_Influenza,Fecha_Td,Fecha_TdSegunda,Fecha_TdRefuerzo,Fecha_TDPA,FPP_USG,FechaEvento) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.NoExpediente, c.NombreCompleto, c.FechaNacimiento, c.Curp, c.Telefono, c.Direccion, c.Gestas, c.Paras, c.Abortos, c.Cesareas, c.Emigro, c.ConsultaPregestacional, c.FechaUltimoEvento, c.FechaUlmaMenstruacion, c.FechaProbableParto, c.FechaConsulta, c.FechaVacunaTDPrimera, c.FechaVacunaTDSegunda, c.FechaVacunaTDRefuerzo, c.FechaVacunaTDPA, c.FechaVacunaInfluenza, c.FechaProbableUSG, c.FechaEvento)
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
	rows, err := bd.Query(`SELECT idMujer, noExpediente, Nombre, curp, telefono, COALESCE(Edad,''), COALESCE(FechaNacimiento, ''), domicilio_referencia, gestas, paras, aborto, cesareas, emigro, COALESCE(Consulta_RiesgoPreg, ''), COALESCE(Fecha_UltimoParto, ''), COALESCE(Fecha_Ult_Mestruacion,''), COALESCE(PFP,''), COALESCE(Fecha_Consulta,''), COALESCE(Fecha_Influenza,''), COALESCE(Fecha_Td,''), COALESCE(Fecha_TdSegunda,''), COALESCE(Fecha_TdRefuerzo,''), COALESCE(Fecha_TDPA,''), COALESCE(FPP_USG,''), COALESCE(FechaEvento,'')  FROM mujer`)
	if err != nil {
		return embarazadas, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var embarazada types.Embarazada
		err := rows.Scan(&embarazada.Id, &embarazada.NoExpediente, &embarazada.NombreCompleto, &embarazada.Curp, &embarazada.Telefono, &embarazada.Edad, &embarazada.FechaNacimiento, &embarazada.Direccion, &embarazada.Gestas, &embarazada.Paras, &embarazada.Abortos, &embarazada.Cesareas, &embarazada.Emigro, &embarazada.ConsultaPregestacional, &embarazada.FechaUltimoEvento, &embarazada.FechaUlmaMenstruacion, &embarazada.FechaProbableParto, &embarazada.FechaConsulta, &embarazada.FechaVacunaInfluenza, &embarazada.FechaVacunaTDPrimera, &embarazada.FechaVacunaTDSegunda, &embarazada.FechaVacunaTDRefuerzo, &embarazada.FechaVacunaTDPA, &embarazada.FechaProbableUSG, &embarazada.FechaUltimoEvento)
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
	rows, err := bd.Query("SELECT * FROM t_municipio order by municipio asc")
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
	filas, err := bd.Query(`SELECT id_localidad,localidad FROM t_localidad WHERE id_municipio=? order by localidad asc `, id)
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
