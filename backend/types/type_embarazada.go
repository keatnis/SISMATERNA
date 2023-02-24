package embarazada

type Embarazada struct {
	Id                     int64  `json:"id"`
	NoExpediente           string `json:"noExpediente"`
	NombreCompleto         string `json:"nombre"`
	Curp                   string `json:"curp"`
	FechaNacimiento        string `json:"FechaNacimiento,omitempty"`
	Edad                   int64  `json:"edad"`
	Direccion              string `json:"domicilioReferencia"`
	Localidad              string `json:"localidad"`
	Municipio              string `json:"municipio"`
	Telefono               string `json:"telefono,omitempty"`
	LenguaIndigena         string `json:"lenguaIndigena"`
	Emigro                 int64  `json:"emigro"`
	Derechohabiencia       string `json:"derechohabiencia"`
	Violencia              string `json:"violencia"`
	Comorbilidades         string `json:"comorbilidades"`
	AsistenciaPreg         int64  `json:"AsistenciaPreg"`
	ConsultaPregestacional string `json:"ConsultaPregestacional,omitempty"`
	Gestas                 int64  `json:"gestas"`
	Paras                  int64  `json:"paras"`
	Abortos                int64  `json:"abortos"`
	Cesareas               int64  `json:"cesareas"`
	FechaUltimoEvento      string `json:"FechaUltimoEvento,omitempty"`
	PresentoComplicaiones  int64  `json:"presentComplicaciones"`
	Complicaciones         string `json:"complicaciones"`
	FechaUlmaMenstruacion  string `json:"FechaUlmaMenstruacion,omitempty"`
	FechaProbableParto     string `json:"FechaProbableParto,omitempty"`
	SGA                    string `json:"SGA"`
	FechaConsulta          string `json:"FechaConsulta,omitempty"`
	NumConsultasEmbarazo   int64  `json:"noConsultasEmbarazo"`
	NumConsultaMes         int64  `json:"noConsultasMes"`
	Rubeola                int64  `json:"rubeola"`
	FechaVacunaTDPrimera   string `json:"FechaVacunaTDPrimera,omitempty"`
	FechaVacunaTDSegunda   string `json:"FechaVacunaTDSegunda,omitempty"`
	FechaVacunaTDRefuerzo  string `json:"FechaVacunaTDRefuerzo,omitempty"`
	FechaVacunaTDPA        string `json:"FechaVacunaTDPA,omitempty"`
	FechaVacunaInfluenza   string `json:"FechaVacunaInfluenza,omitempty"`
	VacunaCOVID            int64  `json:"vacunaCOVID"`
	GrupoRH                string `json:"grupoRH"`
	EGO                    string `json:"ego"`
	BiometriHematica       string `json:"BiometriaHematica"`
	Leucocitos             string `json:"leucocitos"`
	Plaquetas              string `json:"plaquetas"`
	VDRL                   string `json:"vdrl"`
	VIH                    string `json:"vih"`
	EstadoGlucosa          string `json:"estadoGlucosa"`
	ResultadoGlucosa       string `json:"resultadoGlucosa"`
	CaracteristicasFetls   string `json:"caracteristicasFetls"`
	Malformaciones         string `json:"malformaciones"`
	LiquidoAmiotico        string `json:"LiquidoAmiotico"`
	Placenta               string `json:"placenta"`
	FechaProbableUSG       string `json:"FechaProbableUSG,omitempty"`
	Referencia             string `json:"referencia"`
	AcudioReferencia       string `json:"acudio_refe"`
	Contrareferencia       string `json:"contrareferencia"`
	PlanSeguridad          string `json:"planSeguridad"`
	LugarParto             string `json:"lugarParto"`
	QuienAtenderaParto     string `json:"quienAtenderaParto"`
	TransporteAME          string `json:"transporteAME"`
	FechaEvento            string `json:"FechaEvento,omitempty"`
	Observaciones          string `json:"observaciones"`
	Diagnostico            string `json:"diagnostico"`
}
type Municipio struct {
	Id_municipio    int    `json:"id_municipio"`
	NombreMunicipio string `json:"nombreMunicipio"`
}
type Localidad struct {
	Id_localidad    int    `json:"id_localidad"`
	NombreLocalidad string `json:"NombreLocalidad"`
}
