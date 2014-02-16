package nextbus

type Message struct {
	Id                        int    `xml: "id,attr"`
	Creator                   string `xml: "supervisor,attr"`
	SendToBuses               bool   `xml: "sendToBuses,attr"`
	startBoundary             int    `xml: "startBoundary,attr"`
	startBoundaryStr          string `xml: "startBoundaryStr,attr"`
	endBoundary               int    `xml: "endBoundary,attr"`
	endBoundaryStr            string `xml: "endBoundaryStr,attr"`
	RouteForConfiguredMessage Route
	Text                      string
	PhoneMeText               string
	SMSText                   string
	Priority                  int
}
