package nextbus

type Agency struct {
	Tag         string `xml: "tag,attr"`
	Title       string `xml: "title,attr"`
	RegionTitle string `xml: "regionTitle,attr"`
	ShortTitle  string `xml: "shortTitle,attr"`
}
