package nextbus

type Agency struct {
	Tag         string `xml: "tag,attr"`
	Title       string `xml: "title,attr"`
	RegionTitle string `xml: "regionTitle,attr"`
	ShortTitle  string `xml: "shortTitle,attr"`
}

type Route struct {
	Tag           string  `xml: "tag,attr"`
	Title         string  `xml: "title,attr"`
	ShortTitle    string  `xml: "shortTitle,attr"`
	Color         string  `xml: "color,attr"`
	OppositeColor string  `xml: "oppositeColor,attr"`
	LatMin        float32 `xml: "latMin,attr"`
	LatMax        float32 `xml: "latMax,attr"`
	LonMin        float32 `xml: "lonMin,attr"`
	LonMax        float32 `xml: "lonMax,attr"`
	Stops         []Stop
	Directions    []Direction
	Paths         []Path
	Messages      []Message
}

type Stop struct {
	Tag        string  `xml: "tag,attr"`
	Title      string  `xml: "title,attr"`
	ShortTitle string  `xml: "shortTitle,attr"`
	Lat        float32 `xml: "lat,attr"`
	Lon        float32 `xml: "lon,attr"`
	StopId     int     `xml: "stopId,attr"`
}

type Direction struct {
	Tag            string `xml: "tag,attr"`
	Title          string `xml: "title,attr"`
	Name           string `xml: "name,attr"`
	UseForUI       bool   `xml: "useForUI,attr"`
	Stops          []Stop
	PredictionList []Prediction
	Predictions    Predictions
}

type Path struct {
	Points []Point
}

type Point struct {
	Lat float32 `xml: "lat,attr"`
	Lon float32 `xml: "lon,attr"`
}
