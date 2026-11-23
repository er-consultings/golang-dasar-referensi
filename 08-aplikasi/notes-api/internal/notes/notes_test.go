package notes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPI(t *testing.T) {
	srv := httptest.NewServer(Handler(NewStore()))
	defer srv.Close()

	resp, err := http.Post(srv.URL+"/notes", "application/json", strings.NewReader(`{"title":"Learning Go","body":"slices chapter"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("POST status = %d, want 201", resp.StatusCode)
	}
	var n Note
	if err := json.NewDecoder(resp.Body).Decode(&n); err != nil {
		t.Fatal(err)
	}
	if n.ID != 1 || n.Title != "Learning Go" {
		t.Errorf("note = %+v", n)
	}

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		want   int
	}{
		{"get existing", "GET", "/notes/1", "", 200},
		{"get missing", "GET", "/notes/9", "", 404},
		{"id is not a number", "GET", "/notes/abc", "", 400},
		{"empty title", "POST", "/notes", `{"title":" "}`, 400},
		{"list", "GET", "/notes", "", 200},
		{"delete", "DELETE", "/notes/1", "", 204},
		{"delete again", "DELETE", "/notes/1", "", 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, srv.URL+tt.path, strings.NewReader(tt.body))
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if resp.StatusCode != tt.want {
				t.Errorf("%s %s = %d, want %d", tt.method, tt.path, resp.StatusCode, tt.want)
			}
		})
	}
}
