package types

type Embarazada struct {
	Id                     int    `json:"id"`
	NoExpediente           string `json:"noExpediente"`
	NombreCompleto         string `json:"nombre"`
	Curp                   string `json:"curp"`
	Direccion              string `json:"domicilioReferencia"`
	Telefono               int64  `json:"telefono,string,omitempty"`
	FechaNacimiento        string `json:"FechaNacimiento,omitempty"`
	Gestas                 string `json:"gestas"`
	Paras                  string `json:"paras"`
	Abortos                string `json:"abortos"`
	Cesareas               string `json:"cesareas"`
	DondeMigro             string `json:"dondeEmigro"`
	ConsultaPregestacional string `json:"ConsultaPregestacional,omitempty"`
	FechaUltimoEvento      string `json:"FechaUltimoEvento,omitempty"`
}
type Municipio struct {
	Id_municipio    int    `json:"id_municipio"`
	NombreMunicipio string `json:"nombreMunicipio"`
}
type Localidad struct {
	Id_localidad    int    `json:"id_localidad"`
	NombreLocalidad string `json:"NombreLocalidad"`
}
