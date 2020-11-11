package util

import (
	"net/http/httptest"
	"testing"
)

func TestJSONResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	JSONResponse(rr, 200, struct{ Response string }{Response: "Test"})
	if rr.Body.String() != `{"Response":"Test"}` {
		t.Error("JSONResponse failed")
	}
}
