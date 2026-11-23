# golang-dasar-referensi

Repositori pendamping kursus **Dasar Pemrograman Go untuk Pemula** di LearnFlow. Isinya program yang bisa
langsung dijalankan untuk setiap materi, tiga aplikasi kecil, dan 16 soal latihan bergaya LeetCode
tingkat Easy, semuanya dengan test.

## Mulai

Butuh Go 1.23 atau lebih baru.

```bash
git clone https://github.com/er-consultings/golang-dasar-referensi.git
cd golang-dasar-referensi
go test ./...
```

Semua perintah dijalankan dari akar repositori, tempat `go.mod` berada.

## Isi

| Folder | Isi |
|---|---|
| `01-mengenal-go/` | program pertama |
| `02-dasar-bahasa/` | variabel, tipe data, percabangan, perulangan |
| `03-fungsi-error/` | fungsi, error, defer, panic |
| `04-struktur-data/` | slice, map, struct, pointer, string, mengurutkan |
| `05-interface/` | interface dan komposisi |
| `06-struktur-proyek/kasir/` | contoh proyek dengan `cmd/` dan `internal/` |
| `07-test-dan-konkurensi/` | table-driven test, goroutine, channel, Mutex |
| `08-aplikasi/` | hitung-kata (CLI), belanja-cli (berkas JSON), catatan-api (HTTP) |
| `09-latihan-soal/` | 16 soal Easy bergaya LeetCode, lihat README di dalamnya |
| `latihan/` | tempat Anda mengerjakan soal sendiri |

## Materi dan folder

| Bagian | Materi | Folder | Perintah |
|---|---|---|---|
| Mengenal Go | Program pertama: run, build, dan package main | `01-mengenal-go/halo/` | `go run ./01-mengenal-go/halo Sari` |
| Mengenal Go | Memakai repositori referensi kursus | `.` | `go test ./...` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Variabel, konstanta, dan zero value | `02-dasar-bahasa/variabel/` | `go run ./02-dasar-bahasa/variabel` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Tipe data dasar dan konversi | `02-dasar-bahasa/tipe-data/` | `go run ./02-dasar-bahasa/tipe-data` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Percabangan: if dan switch | `02-dasar-bahasa/percabangan/` | `go run ./02-dasar-bahasa/percabangan` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Perulangan dengan for dan range | `02-dasar-bahasa/perulangan/` | `go run ./02-dasar-bahasa/perulangan` |
| Fungsi dan Penanganan Error | Fungsi: parameter, beberapa nilai kembali, variadic | `03-fungsi-error/fungsi/` | `go run ./03-fungsi-error/fungsi` |
| Fungsi dan Penanganan Error | Error adalah nilai: pola if err != nil | `03-fungsi-error/error/` | `go run ./03-fungsi-error/error` |
| Fungsi dan Penanganan Error | defer, panic, dan recover | `03-fungsi-error/defer-panic/` | `go run ./03-fungsi-error/defer-panic` |
| Struktur Data: Slice, Map, Struct, Pointer | Array dan slice | `04-struktur-data/slice/` | `go run ./04-struktur-data/slice` |
| Struktur Data: Slice, Map, Struct, Pointer | Map: pasangan key dan value | `04-struktur-data/map/` | `go run ./04-struktur-data/map` |
| Struktur Data: Slice, Map, Struct, Pointer | Struct dan method | `04-struktur-data/struct-method/` | `go run ./04-struktur-data/struct-method` |
| Struktur Data: Slice, Map, Struct, Pointer | Pointer tanpa rasa takut | `04-struktur-data/pointer/` | `go run ./04-struktur-data/pointer` |
| Struktur Data: Slice, Map, Struct, Pointer | Mengolah string dan mengurutkan data | `04-struktur-data/string-dan-urut/` | `go run ./04-struktur-data/string-dan-urut` |
| Interface dan Komposisi | Interface: perilaku, bukan warisan | `05-interface/bentuk/` | `go run ./05-interface/bentuk` |
| Interface dan Komposisi | Komposisi dengan embedding | `05-interface/komposisi/` | `go run ./05-interface/komposisi` |
| Modul, Paket, dan Struktur Proyek | Go modules dan dependensi | `go.mod` | `go mod tidy` |
| Modul, Paket, dan Struktur Proyek | Paket, visibilitas, dan penamaan paket | `06-struktur-proyek/kasir/internal/` | `go doc ./06-struktur-proyek/kasir/internal/produk` |
| Modul, Paket, dan Struktur Proyek | Struktur proyek dasar yang rapi | `06-struktur-proyek/kasir/` | `go run ./06-struktur-proyek/kasir/cmd/kasir` |
| Kebiasaan Baik, Test, dan Konkurensi | Gaya kode: gofmt, go vet, penamaan, dan komentar | `.` | `gofmt -l . && go vet ./...` |
| Kebiasaan Baik, Test, dan Konkurensi | Menulis test dengan paket testing | `07-test-dan-konkurensi/diskon/` | `go test -v ./07-test-dan-konkurensi/diskon` |
| Kebiasaan Baik, Test, dan Konkurensi | Menulis test dengan paket testing | `06-struktur-proyek/kasir/internal/pesanan/` | `go test -v ./06-struktur-proyek/kasir/internal/pesanan` |
| Kebiasaan Baik, Test, dan Konkurensi | Pengenalan goroutine dan channel | `07-test-dan-konkurensi/konkurensi/` | `go run -race ./07-test-dan-konkurensi/konkurensi` |
| Membangun Aplikasi Pertama | Aplikasi 1: membaca input, berkas, dan flag | `08-aplikasi/hitung-kata/` | `go run ./08-aplikasi/hitung-kata -top 3 README.md` |
| Membangun Aplikasi Pertama | Aplikasi 2: CLI daftar belanja dengan berkas JSON | `08-aplikasi/belanja-cli/` | `go run ./08-aplikasi/belanja-cli tambah kopi 2 18000` |
| Membangun Aplikasi Pertama | Aplikasi 3: API HTTP pertama dengan net/http | `08-aplikasi/catatan-api/` | `go run ./08-aplikasi/catatan-api` |
| Membangun Aplikasi Pertama | Dari ide ke aplikasi: langkah membuat program sendiri | `08-aplikasi/` | `go test ./08-aplikasi/...` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Membaca soal LeetCode dan menerjemahkannya ke Go | `09-latihan-soal/_template/` | `cp -r 09-latihan-soal/_template latihan/two-sum` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Membaca soal LeetCode dan menerjemahkannya ke Go | `latihan/` | `go test ./latihan/two-sum` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Kompleksitas waktu dan ruang (Big-O) tanpa rumus rumit | `09-latihan-soal/02-two-sum/` | `go test -v ./09-latihan-soal/02-two-sum` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 1: hash map untuk mencari dan menghitung | `09-latihan-soal/02-two-sum/` | `go test ./09-latihan-soal/02-two-sum` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 1: hash map untuk mencari dan menghitung | `09-latihan-soal/03-contains-duplicate/` | `go test ./09-latihan-soal/03-contains-duplicate` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 1: hash map untuk mencari dan menghitung | `09-latihan-soal/04-valid-anagram/` | `go test ./09-latihan-soal/04-valid-anagram` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 1: hash map untuk mencari dan menghitung | `09-latihan-soal/13-roman-to-integer/` | `go test ./09-latihan-soal/13-roman-to-integer` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 2: dua pointer | `09-latihan-soal/05-valid-palindrome/` | `go test ./09-latihan-soal/05-valid-palindrome` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 2: dua pointer | `09-latihan-soal/06-reverse-string/` | `go test ./09-latihan-soal/06-reverse-string` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 2: dua pointer | `09-latihan-soal/07-merge-sorted-array/` | `go test ./09-latihan-soal/07-merge-sorted-array` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 2: dua pointer | `09-latihan-soal/11-move-zeroes/` | `go test ./09-latihan-soal/11-move-zeroes` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 3: stack dan binary search | `09-latihan-soal/08-valid-parentheses/` | `go test ./09-latihan-soal/08-valid-parentheses` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 3: stack dan binary search | `09-latihan-soal/09-binary-search/` | `go test ./09-latihan-soal/09-binary-search` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 4: satu kali jalan dan pemrograman dinamis sederhana | `09-latihan-soal/01-fizz-buzz/` | `go test ./09-latihan-soal/01-fizz-buzz` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 4: satu kali jalan dan pemrograman dinamis sederhana | `09-latihan-soal/10-best-time-to-buy-and-sell-stock/` | `go test ./09-latihan-soal/10-best-time-to-buy-and-sell-stock` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 4: satu kali jalan dan pemrograman dinamis sederhana | `09-latihan-soal/12-climbing-stairs/` | `go test -v ./09-latihan-soal/12-climbing-stairs` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 5: linked list dan binary tree | `09-latihan-soal/14-reverse-linked-list/` | `go test ./09-latihan-soal/14-reverse-linked-list` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 5: linked list dan binary tree | `09-latihan-soal/15-merge-two-sorted-lists/` | `go test ./09-latihan-soal/15-merge-two-sorted-lists` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 5: linked list dan binary tree | `09-latihan-soal/16-maximum-depth-of-binary-tree/` | `go test ./09-latihan-soal/16-maximum-depth-of-binary-tree` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Rencana latihan dan cara tidak buntu | `09-latihan-soal/` | `go test ./09-latihan-soal/...` |
| Penutup dan Uji Pengetahuan Akhir | Rangkuman dan langkah berikutnya | `.` | `go test ./...` |

## Perintah yang sering dipakai

```bash
go run ./02-dasar-bahasa/variabel   # jalankan satu contoh
go test ./...                       # semua test
go test -race ./...                 # dengan detektor data race
gofmt -l .                          # berkas yang formatnya belum rapi
go vet ./...                        # pemeriksaan kesalahan umum
```
