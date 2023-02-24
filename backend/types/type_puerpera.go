package embarazada

type Puerpera struct {
	Consultasiete    string `json:"consultasiete,omitempty"`
	Consultaveinte   string `json:"consultaveinte,omitempty"`
	Consultacuarenta string `json:"consultacuarenta,omitempty"`
	Signos           string `json:"signos"`
	Atencionparto    string `json:"atencionparto"`
	Lugarparto       string `json:"lugarparto"`
	Resolucion       string `json:"resolucion"`
	Producto         string `json:"producto"`
	Apeo             string `json:"apeo "`
	Aceptante        string `json:"aceptante"`
	Pregestacional   int64  `json:"pregestacional"`
}
