package controller_embarazada

import (
	db "backendmod/database"
	embarazada "backendmod/types"
)

func InsertEmbarazada(c embarazada.Embarazada) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	sentenciaPreparada, err := db.Prepare("INSERT INTO mujer (curp, noexpediente, nombre, fechanacimiento, domicilio_referencia, localidad, municipio, telefono, lengua_indig, emigro, derechohabiencia, violencia, consulta_riesgopreg, comorbilidades, asistencia_preg , gestas, paras, abortos, cesareas, fecha_ultimoparto,presento_comp, tipo_comp, fecha_ult_mestruacion, fpp, sdg_actual, fecha_consulta, no_consultaembarazo, no_consultames, rubeola, fecha_td, fecha_tdsegunda, fecha_tdrefuerzo, fecha_tdpa, fecha_influenza, covid, grupo_rh, ego, biometria_ematica, leucocitos, plaquetas, vdrl, vih, glucosa_estado, glucosa_resultado, caracteristicas_fetls, malformaciones, liquido_amniotico, placenta, fpp_usg, referencia, acudio_refe, contrareferencia, plan_seguridad, donde_atenderaparto, quien_atenderaparto)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Curp, c.NoExpediente, c.NombreCompleto, c.FechaNacimiento, c.Direccion, c.Localidad, c.Municipio, c.Telefono, c.LenguaIndigena, c.Emigro, c.Derechohabiencia, c.Violencia, c.ConsultaPregestacional, c.Comorbilidades, c.AsistenciaPreg, c.Gestas, c.Paras, c.Abortos, c.Cesareas, c.FechaUltimoEvento, c.PresentoComplicaiones, c.Complicaciones, c.FechaUlmaMenstruacion, c.FechaProbableParto, c.SGA, c.FechaConsulta, c.NumConsultasEmbarazo, c.NumConsultaMes, c.Rubeola, c.FechaVacunaTDPrimera, c.FechaVacunaTDSegunda, c.FechaVacunaTDRefuerzo, c.FechaVacunaTDPA, c.FechaVacunaInfluenza, c.VacunaCOVID, c.GrupoRH, c.EGO, c.BiometriHematica, c.Leucocitos, c.Plaquetas, c.VDRL, c.VIH, c.EstadoGlucosa, c.ResultadoGlucosa, c.CaracteristicasFetls, c.Malformaciones, c.LiquidoAmiotico, c.Placenta, c.FechaProbableUSG, c.Referencia, c.AcudioReferencia, c.Contrareferencia, c.PlanSeguridad, c.LugarParto, c.QuienAtenderaParto)
	if err != nil {
		return err
	}
	return err

	/*

		_, err = db.Exec("INSERT INTO mujer (curp, noexpediente, nombre, fechanacimiento, domicilio_referencia, localidad, municipio, telefono, lengua_indig, emigro, derechohabiencia, detenciones, consulta_riesgopreg, comorbilidades, asistencia_preg , gestas, paras, abortos, cesareas, fecha_ultimoparto, tipo_comp, fecha_ult_mestruacion, fpp, sdg_actual, fecha_consulta, no_consultaembarazo, no_consultames, rubeola, fecha_td, fecha_tdsegunda, fecha_tdrefuerzo, fecha_tdpa, fecha_influenza, covid, grupo_rh, ego, biometria_ematica, leucocitos, plaquetas, vdrl, vih, glucosa_estado, glucosa_resultado, caracteristicas_fetls, malformaciones, liquido_amniotico, placenta, fpp_usg, referencia, acudio_refe, contrareferencia, plan_seguridad, signos_alarma, donde_atenderaparto, quien_atenderaparto, ame, fechaevento, observaciones, diagnostico)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			c.Curp, c.NoExpediente, c.NombreCompleto, c.FechaNacimiento, c.Direccion, c.Localidad, c.Municipio, c.Telefono, c.LenguaIndigena, c.Emigro, c.Derechohabiencia, c.Detenciones, c.ConsultaPregestacional, c.Comorbilidades, c.AsistenciaPreg, c.Gestas, c.Paras, c.Abortos, c.Cesareas, c.FechaUltimoEvento, c.Complicaciones, c.FechaUlmaMenstruacion, c.FechaProbableParto, c.SGA, c.FechaConsulta, c.NumConsultasEmbarazo, c.NumConsultaMes, c.Rubeola, c.FechaVacunaTDPrimera, c.FechaVacunaTDSegunda, c.FechaVacunaTDRefuerzo, c.FechaVacunaTDPA, c.FechaVacunaInfluenza, c.VacunaCOVID, c.GrupoRH, c.EGO, c.BiometriHematica, c.Leucocitos, c.Plaquetas, c.VDRL, c.VIH, c.EstadoGlucosa, c.ResultadoGlucosa, c.CaracteristicasFetls, c.Malformaciones, c.LiquidoAmiotico, c.Placenta, c.FechaProbableUSG, c.Referencia, c.AcudioReferencia, c.Contrareferencia, c.PlanSeguridad, c.SignosAlarma, c.LugarParto, c.QuienAtenderaParto, c.TransporteAME, c, c.FechaEvento, c.Observaciones, c.Diagnostico)
		if err != nil {
			return err
		}
		return nil
	*/
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
func updateEmbarazada(e embarazada.Embarazada) error {
	bd, err := db.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE video_games SET name = ?, genre = ?, year = ? WHERE id = ?",
		e.NombreCompleto, e.Direccion)
	return err
}
func GetEmbarazada() ([]embarazada.Embarazada, error) {
	//Declare an array because if there's error, we return it empty
	embarazadas := []embarazada.Embarazada{}
	bd, err := db.GetDB()
	if err != nil {
		return embarazadas, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query(`SELECT idMujer, curp, noexpediente, nombre, COALESCE(fechanacimiento,''), COALESCE(Edad,''),  domicilio_referencia, localidad, municipio, telefono, lengua_indig, emigro, derechohabiencia, violencia, consulta_riesgopreg, comorbilidades, gestas, paras, abortos, cesareas, COALESCE(fecha_ultimoparto, ''), presento_comp, tipo_comp, COALESCE(fecha_ult_mestruacion,''), COALESCE(fpp,''), sdg_actual, COALESCE(fecha_consulta,''), no_consultaembarazo, no_consultames, rubeola, fecha_td, fecha_tdsegunda, fecha_tdrefuerzo, fecha_tdpa,  COALESCE(fecha_influenza,''), covid, grupo_rh, ego,biometria_ematica, leucocitos, plaquetas, vdrl, vih, glucosa_estado, glucosa_resultado, caracteristicas_fetls, malformaciones,liquido_amniotico, placenta, COALESCE(fpp_usg,''), referencia, acudio_refe, contrareferencia, plan_seguridad, donde_atenderaparto,quien_atenderaparto FROM mujer`)
	if err != nil {
		return embarazadas, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var embarazada embarazada.Embarazada
		err := rows.Scan(&embarazada.Id, &embarazada.Curp, &embarazada.NoExpediente, &embarazada.NombreCompleto, &embarazada.FechaNacimiento, &embarazada.Edad, &embarazada.Direccion, &embarazada.Localidad, &embarazada.Municipio, &embarazada.Telefono, &embarazada.LenguaIndigena, &embarazada.Emigro, &embarazada.Derechohabiencia, &embarazada.Violencia, &embarazada.ConsultaPregestacional, &embarazada.Comorbilidades, &embarazada.Gestas, &embarazada.Paras, &embarazada.Abortos, &embarazada.Cesareas, &embarazada.FechaUltimoEvento, &embarazada.PresentoComplicaiones, &embarazada.Complicaciones, &embarazada.FechaUlmaMenstruacion, &embarazada.FechaProbableParto, &embarazada.SGA, &embarazada.FechaConsulta, &embarazada.NumConsultasEmbarazo, &embarazada.NumConsultaMes, &embarazada.Rubeola, &embarazada.FechaVacunaTDPrimera, &embarazada.FechaVacunaTDSegunda, &embarazada.FechaVacunaTDRefuerzo, &embarazada.FechaVacunaTDPA, &embarazada.FechaVacunaInfluenza, &embarazada.VacunaCOVID, &embarazada.GrupoRH, &embarazada.EGO, &embarazada.BiometriHematica, &embarazada.Leucocitos, &embarazada.Plaquetas, &embarazada.VDRL, &embarazada.VIH, &embarazada.EstadoGlucosa, &embarazada.ResultadoGlucosa, &embarazada.CaracteristicasFetls, &embarazada.Malformaciones, &embarazada.LiquidoAmiotico, &embarazada.Placenta, &embarazada.FechaProbableUSG, &embarazada.Referencia, &embarazada.AcudioReferencia, &embarazada.Contrareferencia, &embarazada.PlanSeguridad, &embarazada.LugarParto, &embarazada.QuienAtenderaParto)
		if err != nil {
			return embarazadas, err
		}
		// and append it to the array
		embarazadas = append(embarazadas, embarazada)
	}
	return embarazadas, nil
}

func GetMunicipios() ([]embarazada.Municipio, error) {
	municipios := []embarazada.Municipio{}
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
		var municipio embarazada.Municipio
		err := rows.Scan(&municipio.Id_municipio, &municipio.NombreMunicipio)
		if err != nil {
			return municipios, err
		}
		municipios = append(municipios, municipio)
	}
	return municipios, nil

}

func ObtenerLocalidades(id int64) ([]embarazada.Localidad, error) {
	localidades := []embarazada.Localidad{}
	bd, err := db.GetDB()

	if err != nil {
		return localidades, err
	}

	filas, err := bd.Query(`SELECT id_localidad,localidad FROM t_localidad WHERE id_municipio=? order by localidad asc `, id)
	if err != nil {
		return localidades, err
	}
	defer bd.Close()

	var localidad embarazada.Localidad
	for filas.Next() {
		err := filas.Scan(&localidad.Id_localidad, &localidad.NombreLocalidad)

		if err != nil {
			return localidades, err
		}
		localidades = append(localidades, localidad)
	}
	defer filas.Close()
	return localidades, nil
}
