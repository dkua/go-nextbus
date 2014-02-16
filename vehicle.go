package nextbus

type Vehicle struct {
	Id              string  `xml: "id,attr"`
	RouteTag        string  `xml: "routeTag,attr"`
	DirTag          string  `xml: "dirTag,attr"`
	Lat             float32 `xml: "lat,attr"`
	Lon             float32 `xml: "lon,attr"`
	SecsSinceReport int     `xml: "secsSinceReport,attr"`
	Predictable     bool    `xml: "predictable,attr"`
	Heading         int     `xml: "heading,attr"`
	SpeedKmHr       float32 `xml: "speedKmHr,attr"`
}

type LastTime struct {
	Time int `xml: "time,attr"`
}
