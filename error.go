package nextbus

type Error struct {
	ShouldRetry bool   `xml:"shouldRetry,attr"`
	Message     string `xml:"Error"`
}
