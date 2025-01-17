package dao // data access objects

type Polygon struct {
	Geometry interface{} `json:"geometry"`
	Type     string      `json:"type"`
	Color    string      `json:"color"`
}

func (p *Polygon) SetColor(color string) {
	p.Color = color

}

//editable, clip, visible polycount

//home/userone/Documentos/AppReact/flaskAPI/zonif/zonifModels.py

// Compartida por Vra y Zonif
type Zone struct {
	Id        string    `json:"id"`
	ZonifId   string    `json:"zonifid"`
	Name      string    `json:"name"`
	Editable  bool      `json:"editable"`
	Clip      string    `json:"clip"`
	Visible   bool      `json:"visible"`
	Area      string    `json:"area"`
	MaxValue  string    `json:"maxvalue"`
	MinValue  string    `json:"minvalue"`
	Polycount int       `json:"polycount"`
	Polygons  []Polygon `json:"polygons"`
	Color     string    `json:"color"`
}

func (z *Zone) SetColor(color string) {
	z.Color = color

}
