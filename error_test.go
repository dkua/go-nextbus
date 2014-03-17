package nextbus

import (
	"encoding/xml"
	"testing"
)

const TestErrorXML = `<body><Error shouldRetry="true">Agency server cannot accept client while status is: agency name = sf-muni,status = UNINITIALIZED, client count = 0, last even t = 0 seconds ago Could not get route list for agency tag "sf-muni". Either the route tag is bad or the system is initializing.</Error></body>`

func setUp() *Error {
	content := []byte(TestErrorXML)
	e := new(Error)
	xml.Unmarshal(content, e)
	return e
}

func TestError(t *testing.T) {
	e := setUp()
	result := e.Message
	expected := `Agency server cannot accept client while status is: agency name = sf-muni,status = UNINITIALIZED, client count = 0, last even t = 0 seconds ago Could not get route list for agency tag "sf-muni". Either the route tag is bad or the system is initializing.`
	if result != expected {
		t.Errorf("Expected \"%v\" in Error tag got \"%v\" instead.", expected, result)
	}
}

func TestShouldRetry(t *testing.T) {
	e := setUp()
	result := e.ShouldRetry
	expected := true
	if result != expected {
		t.Errorf("Expected \"%v\" in shouldRetry attr got \"%v\" instead.", expected, result)
	}
}
