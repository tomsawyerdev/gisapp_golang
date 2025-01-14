package dao // data access objects

type Polygon struct {
	Geometry interface{} `json:"geometry"`
	Type     string      `json:"type"`
	Color    string      `json:"color"`
}

func (p *Polygon) SetColor(color string) {
	p.Color = color

}

type Zone struct {
	Id       int       `json:"id"`
	Area     string    `json:"area"`
	MaxValue string    `json:"maxvalue"`
	MinValue string    `json:"minvalue"`
	Name     string    `json:"name"`
	Polygons []Polygon `json:"polygons"`
	Color    string    `json:"color"`
}

func (z *Zone) SetColor(color string) {
	z.Color = color

}
