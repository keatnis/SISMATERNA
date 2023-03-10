package embarazada

type Usuario struct {
	Id                  int64  `json:"id"`
	Clues               string `json:"clues"`
	Nombreunidad        string `json:"nombreunidad"`
	Usuario             string `json:"usuario"`
	Contraseña          string `json:"contraseña"`
	Confirmarcontraseña string `json:"confirmarcontraseña"`
}
