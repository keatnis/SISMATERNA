package main

type Embarazada struct {
	Id              int    `json:"id"`
	NoExpediente    string `json:"noExpediente"`
	Nombre          string `json:"nombre"`
	Curp            string `json:"curp"`
	Telefono        int64  `json:"telefono,string,omitempty"`
	FechaNacimiento string `json:"FechaNacimiento,omitempty"`
}

func insertarEmbarazada(c Embarazada) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO mujer (No_Expediente, Nombre,FechaNacimiento, curp,Telefono) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.NoExpediente, c.Nombre, c.FechaNacimiento, c.Curp, c.Telefono)
	if err != nil {
		return err
	}
	return nil
}

func establecerFechaSalida(IdVehiculo int64, FechaSalida string) error {
	bd, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE vehiculos SET fecha_salida = ? WHERE id = ?", FechaSalida, IdVehiculo)
	if err != nil {
		return err
	}
	return nil
}

func obtenerVehiculos() ([]Embarazada, error) {
	vehiculos := []Embarazada{}
	bd, err := obtenerBaseDeDatos()

	if err != nil {
		return vehiculos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`SELECT id_mujer,No_Expediente, nombre,curp,telefono, COALESCE(FechaNacimiento, '') FROM mujer`)
	if err != nil {
		return vehiculos, err
	}
	defer filas.Close()
	var vehiculo Embarazada
	for filas.Next() {
		err := filas.Scan(&vehiculo.Id, &vehiculo.NoExpediente, &vehiculo.Nombre, &vehiculo.Curp, &vehiculo.Telefono, &vehiculo.FechaNacimiento)

		if err != nil {
			return vehiculos, err
		}
		vehiculos = append(vehiculos, vehiculo)

	}
	return vehiculos, nil
}
