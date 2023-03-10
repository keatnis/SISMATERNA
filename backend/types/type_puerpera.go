package embarazada

type Puerpera struct {
	Id               int64   `json:"id"`
	Consultasiete    *string `json:"consultasiete,omitempty"`
	Consultaveinte   *string `json:"consultaveinte,omitempty"`
	Consultacuarenta *string `json:"consultacuarenta,omitempty"`
	Signos           *string `json:"signos,omitempty"`
	Atencionparto    *string `json:"atencionparto,omitempty"`
	Lugarparto       *string `json:"lugarparto,omitempty"`
	Resolucion       *string `json:"resolucion,omitempty"`
	Producto         *string `json:"producto,omitempty"`
	Apeo             *string `json:"apeo,omitempty"`
	Aceptante        *string `json:"aceptante,omitempty"`
	Pregestacional   int64   `json:"pregestacional,omitempty"`
}
