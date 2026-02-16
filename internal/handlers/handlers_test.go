package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTest = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{
			key:   "start",
			value: "2026-01-01",
		},
		{
			key:   "end",
			value: "2026-01-02",
		},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postData{
		{
			key:   "start",
			value: "2026-01-01",
		},
		{
			key:   "end",
			value: "2026-01-02",
		},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		{
			key:   "first_name",
			value: "Jairo",
		},
		{
			key:   "last_name",
			value: "Cubillos",
		},
		{
			key:   "email",
			value: "jc@test.com",
		},
		{
			key:   "phone",
			value: "555-555-5555",
		},
	}, http.StatusSeeOther},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTest {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s got status code %d, want %d", e.name, resp.StatusCode, e.expectedStatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			// Create a client that doesn't follow redirects
			client := ts.Client()
			client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}

			resp, err := client.PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s got status code %d, want %d", e.name, resp.StatusCode, e.expectedStatusCode)
			}
		}
	}
}
