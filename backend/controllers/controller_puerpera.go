package controller_embarazada

import (
	db "backendmod/database"
	"backendmod/types"
)

func InsertPuerpera(c types.Puerpera) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	sentenciaPreparada, err := db.Prepare("INSERT INTO puerpera (consulta_siete, consulta_veintiocho, consulta_cuarenta) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Consultasiete, c.Consultaveinte, c.Consultacuarenta)
	if err != nil {
		return err
	}
	return nil
}
func GetPuerpera() ([]types.Puerpera, error) {
	//Declare an array because if there's error, we return it empty
	puerpera := []types.Puerpera{}
	bd, err := db.GetDB()
	if err != nil {
		return puerpera, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query(`SELECT idMujer, noExpediente, Nombre, curp, telefono, COALESCE(Edad,''), COALESCE(FechaNacimiento, ''), domicilio_referencia, gestas, paras, aborto, cesareas, emigro, COALESCE(Consulta_RiesgoPreg, ''), COALESCE(Fecha_UltimoParto, ''), COALESCE(Fecha_Ult_Mestruacion,''), COALESCE(PFP,''), COALESCE(Fecha_Consulta,''), COALESCE(Fecha_Influenza,''), COALESCE(Fecha_Td,''), COALESCE(Fecha_TdSegunda,''), COALESCE(Fecha_TdRefuerzo,''), COALESCE(Fecha_TDPA,''), COALESCE(FPP_USG,''), COALESCE(FechaEvento,'')  FROM mujer`)
	if err != nil {
		return puerpera, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var embarazada types.Embarazada
		err := rows.Scan(&embarazada.Id, &embarazada.NoExpediente, &embarazada.NombreCompleto, &embarazada.Curp, &embarazada.Telefono, &embarazada.Edad, &embarazada.FechaNacimiento, &embarazada.Direccion, &embarazada.Gestas, &embarazada.Paras, &embarazada.Abortos, &embarazada.Cesareas, &embarazada.Emigro, &embarazada.ConsultaPregestacional, &embarazada.FechaUltimoEvento, &embarazada.FechaUlmaMenstruacion, &embarazada.FechaProbableParto, &embarazada.FechaConsulta, &embarazada.FechaVacunaInfluenza, &embarazada.FechaVacunaTDPrimera, &embarazada.FechaVacunaTDSegunda, &embarazada.FechaVacunaTDRefuerzo, &embarazada.FechaVacunaTDPA, &embarazada.FechaProbableUSG, &embarazada.FechaUltimoEvento)
		if err != nil {
			return puerpera, err
		}
		// and append it to the array
		puerpera = append(puerpera, puerpera...)
	}
	return puerpera, nil
}
