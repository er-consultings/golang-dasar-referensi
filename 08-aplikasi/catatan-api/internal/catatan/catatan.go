// Package catatan menyimpan catatan di memori dan melayaninya lewat HTTP JSON.
package catatan

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ErrTidakDitemukan dikembalikan bila ID catatan tidak ada.
var ErrTidakDitemukan = errors.New("catatan tidak ditemukan")

// Catatan adalah satu catatan.
type Catatan struct {
	ID     int       `json:"id"`
	Judul  string    `json:"judul"`
	Isi    string    `json:"isi"`
	Dibuat time.Time `json:"dibuat"`
}

// Penyimpan menyimpan catatan di memori. Aman dipakai dari banyak goroutine,
// karena server HTTP melayani setiap permintaan di goroutine sendiri.
type Penyimpan struct {
	mu     sync.Mutex
	data   map[int]Catatan
	nextID int
}

// NewPenyimpan membuat penyimpan kosong.
func NewPenyimpan() *Penyimpan {
	return &Penyimpan{data: map[int]Catatan{}, nextID: 1}
}

// Tambah menyimpan catatan baru dan mengembalikannya lengkap dengan ID.
func (p *Penyimpan) Tambah(judul, isi string) Catatan {
	p.mu.Lock()
	defer p.mu.Unlock()
	c := Catatan{ID: p.nextID, Judul: judul, Isi: isi, Dibuat: time.Now().UTC()}
	p.data[c.ID] = c
	p.nextID++
	return c
}

// Semua mengembalikan semua catatan, urut ID.
func (p *Penyimpan) Semua() []Catatan {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make([]Catatan, 0, len(p.data))
	for _, c := range p.data {
		out = append(out, c)
	}
	slices.SortFunc(out, func(a, b Catatan) int { return a.ID - b.ID })
	return out
}

// Ambil mengembalikan satu catatan.
func (p *Penyimpan) Ambil(id int) (Catatan, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	c, ok := p.data[id]
	if !ok {
		return Catatan{}, ErrTidakDitemukan
	}
	return c, nil
}

// Hapus membuang satu catatan.
func (p *Penyimpan) Hapus(id int) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.data[id]; !ok {
		return ErrTidakDitemukan
	}
	delete(p.data, id)
	return nil
}

// Handler mengembalikan router HTTP untuk penyimpan ini.
//
//	GET    /catatan       semua catatan
//	POST   /catatan       {"judul": "...", "isi": "..."}
//	GET    /catatan/{id}  satu catatan
//	DELETE /catatan/{id}  hapus catatan
func Handler(p *Penyimpan) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /catatan", func(w http.ResponseWriter, r *http.Request) {
		tulisJSON(w, http.StatusOK, p.Semua())
	})
	mux.HandleFunc("POST /catatan", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Judul string `json:"judul"`
			Isi   string `json:"isi"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			tulisError(w, http.StatusBadRequest, "body harus JSON")
			return
		}
		if strings.TrimSpace(in.Judul) == "" {
			tulisError(w, http.StatusBadRequest, "judul wajib diisi")
			return
		}
		tulisJSON(w, http.StatusCreated, p.Tambah(in.Judul, in.Isi))
	})
	mux.HandleFunc("GET /catatan/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			tulisError(w, http.StatusBadRequest, "id harus angka")
			return
		}
		c, err := p.Ambil(id)
		if errors.Is(err, ErrTidakDitemukan) {
			tulisError(w, http.StatusNotFound, err.Error())
			return
		}
		tulisJSON(w, http.StatusOK, c)
	})
	mux.HandleFunc("DELETE /catatan/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			tulisError(w, http.StatusBadRequest, "id harus angka")
			return
		}
		if err := p.Hapus(id); err != nil {
			tulisError(w, http.StatusNotFound, err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
	return mux
}

func tulisJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func tulisError(w http.ResponseWriter, status int, pesan string) {
	tulisJSON(w, status, map[string]string{"error": pesan})
}
