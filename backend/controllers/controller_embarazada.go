package controller_embarazada

import (
	db "backendmod/database"
	"backendmod/types"
)

func InsertEmbarazada(c types.Embarazada) error {
	/* db, err := db.GetDB()
	if err != nil {
		return err
	}
	*/
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()

	//_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", embarazada.NombreCompleto)
	//return err
	// Preparamos para prevenir inyecciones SQL
	_, err = db.Exec("INSERT INTO mujer `curp`, `noexpediente`, `nombre`, `fechanacimiento`, `domicilio_referencia`, `localidad`, `municipio`, `telefono`, `lengua_indig`, `emigro`, `derechohabiencia`, `detenciones`, `consulta_riesgopreg`, `comorbilidades`, `gestas`, `paras`, `abortos`, `cesareas`, `fecha_ultimoparto`, `tipo_comp`, `fecha_ult_mestruacion`, `fpp`, `sdg_actual`, `fecha_consulta`, `no_consultaembarazo`, `no_consultames`, `rubeola`, `fecha_td`, `fecha_tdsegunda`, `fecha_tdrefuerzo`, `fecha_tdpa`, `fecha_influenza`, `covid`, `grupo_rh`, `ego`, `biometria_ematica`, `leucocitos`, `plaquetas`, `vdrl`, `vih`, `glucosa_estado`, `glucosa_resultado`, `caracteristicas_fetls`, `malformaciones`, `liquido_amniotico`, `placenta`, `fpp_usg`, `referencia`, `acudio_refe`, `contrareferencia`, `plan_seguridad`, `signos_alarma`, `donde_atenderaparto`, `quien_atenderaparto`, `ame`, `fechaevento`, `observaciones`, `diagnostico`)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		c.Curp, c.NoExpediente, c.NombreCompleto, c.FechaNacimiento, c.Direccion, c.Localidad, c.Municipio, c.Telefono, c.LenguaIndigena, c.Emigro, c.Derechohabiencia, c.Detenciones, c.ConsultaPregestacional, c.Comorbilidades, c.Gestas, c.Paras, c.Abortos, c.Cesareas, c.FechaUltimoEvento, c.Complicaciones, c.FechaUlmaMenstruacion, c.FechaProbableParto, c.SGA, c.FechaConsulta, c.NumConsultasEmbarazo, c.NumConsultaMes, c.Rubeola, c.FechaVacunaTDPrimera, c.FechaVacunaTDSegunda, c.FechaVacunaTDRefuerzo, c.FechaVacunaTDPA, c.FechaVacunaInfluenza, c.VacunaCOVID, c.GrupoRH, c.EGO, c.BiometriHematica, c.Leucocitos, c.Plaquetas, c.VDRL, c.VIH, c.EstadoGlucosa, c.ResultadoGlucosa, c.CaracteristicasFetls, c.Malformaciones, c.LiquidoAmiotico, c.Placenta, c.FechaProbableUSG, c.Referencia, c.AcudioReferencia, c.Contrareferencia, c.PlanSeguridad, c.SignosAlarma, c.LugarParto, c.QuienAtenderaParto, c.TransporteAME, c, c.FechaEvento, c.Observaciones, c.Diagnostico)
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
