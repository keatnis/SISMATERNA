package types

type Embarazada struct {
	Id                     int    `json:"id"`
	NoExpediente           string `json:"noExpediente"`
	NombreCompleto         string `json:"nombre"`
	Curp                   string `json:"curp"`
	Direccion              string `json:"domicilioReferencia"`
	Telefono               int64  `json:"telefono,omitempty"`
	FechaNacimiento        string `json:"FechaNacimiento,omitempty"`
	Edad                   string `json:"edad"`
	Gestas                 string `json:"gestas"`
	Paras                  string `json:"paras"`
	Abortos                string `json:"abortos"`
	Cesareas               string `json:"cesareas"`
	Emigro                 string `json:"emigro"`
	ConsultaPregestacional string `json:"ConsultaPregestacional,omitempty"`
	FechaUltimoEvento      string `json:"FechaUltimoEvento,omitempty"`
	FechaUlmaMenstruacion  string `json:"FechaUlmaMenstruacion,omitempty"`
	FechaProbableParto     string `json:"FechaProbableParto,omitempty"`
	FechaConsulta          string `json:"FechaConsulta,omitempty"`
	FechaVacunaTDPrimera   string `json:"FechaVacunaTDPrimera,omitempty"`
	FechaVacunaTDSegunda   string `json:"FechaVacunaTDSegunda,omitempty"`
	FechaVacunaTDRefuerzo  string `json:"FechaVacunaTDRefuerzo,omitempty"`
	FechaVacunaTDPA        string `json:"FechaVacunaTDPA,omitempty"`
	FechaVacunaInfluenza   string `json:"FechaVacunaInfluenza,omitempty"`
	FechaProbableUSG       string `json:"FechaProbableUSG,omitempty"`
	FechaEvento            string `json:"FechaEvento,omitempty"`
}
type Municipio struct {
	Id_municipio    int    `json:"id_municipio"`
	NombreMunicipio string `json:"nombreMunicipio"`
}
type Localidad struct {
	Id_localidad    int    `json:"id_localidad"`
	NombreLocalidad string `json:"NombreLocalidad"`
}
