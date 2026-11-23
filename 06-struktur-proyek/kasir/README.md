# kasir: contoh struktur proyek

```
kasir/
├── cmd/kasir/main.go        program yang dijalankan, hanya merangkai
└── internal/
    ├── produk/              katalog produk (domain)
    └── pesanan/             membuat pesanan, butuh katalog lewat interface kecil
```

- `main` tipis memakai pola `run() error`.
- `pesanan` mendefinisikan interface `Katalog` sendiri, hanya berisi method yang ia pakai.
- `internal/` membuat paket ini tidak bisa diimpor dari luar folder `kasir`.

```bash
go run ./06-struktur-proyek/kasir/cmd/kasir
go test ./06-struktur-proyek/...
```
