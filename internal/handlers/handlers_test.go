package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vadim-vep/booking/internal/models"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"generals", "/generals-quarters", "GET", http.StatusOK},
	{"majors", "/majors-suites", "GET", http.StatusOK},
	{"search", "/search-availability", "GET", http.StatusOK},
	{"contacts", "/contacts", "GET", http.StatusOK},
	{"reservation", "/make-reservation", "GET", http.StatusOK},
	//{"post-search-availability", "/search-availability", "POST", []postData{
	//	{key: "start", value: "2020-01-01"},
	//	{key: "end", value: "2020-01-01"},
	//}, http.StatusOK},
	//{"post-search-availability-JSON", "/search-availability-json", "POST", []postData{
	//	{key: "start", value: "2020-01-01"},
	//	{key: "end", value: "2020-01-01"},
	//}, http.StatusOK},
	//{"make-reservation", "/make-reservation", "POST", []postData{
	//	{key: "first_name", value: "Vad"},
	//	{key: "last_name", value: "Vep"},
	//	{key: "email", value: "vad.vep@g.com"},
	//	{key: "phone", value: "555-555-551"},
	//}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("For %s Expected status code %d, got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func TestRepository_reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, rr.Code)
	}

	//test case when reservation is not in session
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, rr.Code)
	}
	//test with non-existent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Expected response code %d, got %d", http.StatusOK, rr.Code)
	}

}

func TestRepository_PostReservation(t *testing.T) {

}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
