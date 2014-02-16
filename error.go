package nextbus

type Error struct {
	ShouldRetry bool   `xml: "shouldRetry,attr"`
	Error       string `xml: "Error"`
}
