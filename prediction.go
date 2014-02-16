package nextbus

type Predictions struct {
	AgencyTitle                  string `xml: "agencyTitle,attr"`
	RouteTag                     string `xml: "routeTag,attr"`
	RouteTitle                   string `xml: "routeTitle,attr"`
	StopTitle                    string `xml: "stopTitle,attr"`
	DirTitleBecauseNoPredictions string `xml: "dirTitleBecauseNoPredictions,attr"`
	Directions                   []Direction
	Message                      Message
	PredictionList               []Prediction
}

type Prediction struct {
	Seconds           int    `xml: "seconds,attr"`
	Minutes           int    `xml: "minutes,attr"`
	EpochTime         int    `xml: "epochTime,attr"`
	IsDeparture       bool   `xml: "isDeparture,attr"`
	Block             string `xml: "block,attr"`
	DirTag            string `xml: "dirTag,attr"`
	TripTag           string `xml: "tripTag,attr"`
	Branch            string `xml: "branch,attr"`
	AffectedByLayover bool   `xml: "affectedByLayover,attr"`
	IsScheduleBased   bool   `xml: "isScheduleBased,attr"`
	Delayed           bool   `xml: "delayed,attr"`
}

type Schedule struct {
	Tag           string `xml: "tag,attr"`
	Title         string `xml: "title,attr"`
	ScheduleClass string `xml: "scheduleClass,attr"`
	ServiceClass  string `xml: "serviceClass,attr"`
	Direction     string `xml: "direction,attr"`
	Header        []Field
	Rows          []Field
}

type Field struct {
	Stop      string `xml: "stop"`
	Tag       string `xml: "tag,attr"`
	EpochTime int    `xml: "epochTime,attr"`
}
