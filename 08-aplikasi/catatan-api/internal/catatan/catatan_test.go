package catatan

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPI(t *testing.T) {
	srv := httptest.NewServer(Handler(NewPenyimpan()))
	defer srv.Close()

	resp, err := http.Post(srv.URL+"/catatan", "application/json", strings.NewReader(`{"judul":"Belajar Go","isi":"bab slice"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("POST status = %d, ingin 201", resp.StatusCode)
	}
	var c Catatan
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		t.Fatal(err)
	}
	if c.ID != 1 || c.Judul != "Belajar Go" {
		t.Errorf("catatan = %+v", c)
	}

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		want   int
	}{
		{"ambil yang ada", "GET", "/catatan/1", "", 200},
		{"ambil yang tidak ada", "GET", "/catatan/9", "", 404},
		{"id bukan angka", "GET", "/catatan/abc", "", 400},
		{"judul kosong", "POST", "/catatan", `{"judul":" "}`, 400},
		{"daftar", "GET", "/catatan", "", 200},
		{"hapus", "DELETE", "/catatan/1", "", 204},
		{"hapus lagi", "DELETE", "/catatan/1", "", 404},
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
				t.Errorf("%s %s = %d, ingin %d", tt.method, tt.path, resp.StatusCode, tt.want)
			}
		})
	}
}
