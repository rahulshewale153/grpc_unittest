package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Index(w http.ResponseWriter, r *http.Request) {
	WriteOKResponse(w, "Record Save Succssfully")
}
func IndexError(w http.ResponseWriter, r *http.Request) {
	WriteErrorResponse(w, http.StatusBadRequest, "Record Not Succssfully")
}
func TestWriteOKResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

}

func TestWriteErrorResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/indexerror", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexError)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
}
