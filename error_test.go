package nextbus

import (
	"encoding/xml"
	"testing"
)

const TestErrorXML = `<body copyright="All data copyright Toronto Transit Commission 2014."><Error shouldRetry="false">stopId=1052315 is not valid for agency=ttc</Error></body>`

func setUp() *Error {
	content := []byte(TestErrorXML)
	e := new(Error)
	xml.Unmarshal(content, e)
	return e
}

func TestError(t *testing.T) {
	e := setUp()
	result := e.Message
	expected := "stopId=1052315 is not valid for agency=ttc"
	if result != expected {
		t.Errorf("Expected \"%v\" in Error tag got \"%v\" instead.", expected, result)
	}
}

func TestShouldRetry(t *testing.T) {
	e := setUp()
	result := e.ShouldRetry
	expected := false
	if result != expected {
		t.Errorf("Expected \"%v\" in shouldRetry attr got \"%v\" instead.", expected, result)
	}
}
