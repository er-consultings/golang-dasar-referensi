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

Semua perintah dijalankan dari folder utama repositori, tempat `go.mod` berada.

Nama folder bab (`01-mengenal-go`, `02-dasar-bahasa`, dan seterusnya) mengikuti nama bagian kursus. Semua kode di dalamnya, termasuk nama folder program, nama package, variabel, komentar, dan pesan, ditulis dalam bahasa Inggris.

## Isi

| Folder | Isi |
|---|---|
| `01-mengenal-go/` | program pertama |
| `02-dasar-bahasa/` | variabel, tipe data, percabangan, perulangan |
| `03-fungsi-error/` | fungsi, error, defer, panic |
| `04-struktur-data/` | slice, map, struct, pointer, string, mengurutkan, memori |
| `05-interface/` | interface dan komposisi |
| `06-struktur-proyek/cashier/` | contoh proyek dengan `cmd/` dan `internal/` |
| `07-test-dan-konkurensi/` | table-driven test, benchmark, goroutine, channel, mutex, context |
| `08-aplikasi/` | word-count (CLI), shopping-cli (file JSON), notes-api (HTTP dengan graceful shutdown) |
| `09-latihan-soal/` | 16 soal Easy bergaya LeetCode, lihat README di dalamnya |
| `latihan/` | tempat Anda mengerjakan soal sendiri |

## Materi dan folder

| Bagian | Materi | Folder | Perintah |
|---|---|---|---|
| Mengenal Go | Program pertama: run, build, dan package main | `01-mengenal-go/hello/` | `go run ./01-mengenal-go/hello Sari` |
| Mengenal Go | Memakai repositori referensi kursus | `.` | `go test ./...` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Variabel, konstanta, dan zero value | `02-dasar-bahasa/variables/` | `go run ./02-dasar-bahasa/variables` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Tipe data dasar dan konversi | `02-dasar-bahasa/data-types/` | `go run ./02-dasar-bahasa/data-types` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Percabangan: if dan switch | `02-dasar-bahasa/branching/` | `go run ./02-dasar-bahasa/branching` |
| Dasar Bahasa: Variabel, Tipe, dan Alur | Perulangan dengan for dan range | `02-dasar-bahasa/loops/` | `go run ./02-dasar-bahasa/loops` |
| Fungsi dan Penanganan Error | Fungsi: parameter, beberapa nilai kembalian, variadic | `03-fungsi-error/functions/` | `go run ./03-fungsi-error/functions` |
| Fungsi dan Penanganan Error | Error adalah nilai: pola if err != nil | `03-fungsi-error/errors/` | `go run ./03-fungsi-error/errors` |
| Fungsi dan Penanganan Error | defer, panic, dan recover | `03-fungsi-error/defer-panic/` | `go run ./03-fungsi-error/defer-panic` |
| Struktur Data dan Memori | Array dan slice | `04-struktur-data/slice/` | `go run ./04-struktur-data/slice` |
| Struktur Data dan Memori | Map: pasangan key dan value | `04-struktur-data/map/` | `go run ./04-struktur-data/map` |
| Struktur Data dan Memori | Struct dan method | `04-struktur-data/struct-method/` | `go run ./04-struktur-data/struct-method` |
| Struktur Data dan Memori | Pointer tanpa rasa takut | `04-struktur-data/pointer/` | `go run ./04-struktur-data/pointer` |
| Struktur Data dan Memori | Mengolah string dan mengurutkan data | `04-struktur-data/strings-and-sorting/` | `go run ./04-struktur-data/strings-and-sorting` |
| Struktur Data dan Memori | Memori: stack, heap, dan garbage collector | `04-struktur-data/memory/` | `go run ./04-struktur-data/memory` |
| Struktur Data dan Memori | Memori: stack, heap, dan garbage collector | `04-struktur-data/memory/` | `go test -bench . -benchmem ./04-struktur-data/memory` |
| Interface dan Komposisi | Interface: perilaku, bukan pewarisan | `05-interface/shapes/` | `go run ./05-interface/shapes` |
| Interface dan Komposisi | Komposisi dengan embedding | `05-interface/composition/` | `go run ./05-interface/composition` |
| Modul, Package, dan Struktur Proyek | Go modules dan dependensi | `go.mod` | `go mod tidy` |
| Modul, Package, dan Struktur Proyek | Package, visibilitas, dan cara menamainya | `06-struktur-proyek/cashier/internal/` | `go doc ./06-struktur-proyek/cashier/internal/product` |
| Modul, Package, dan Struktur Proyek | Struktur proyek dasar yang rapi | `06-struktur-proyek/cashier/` | `go run ./06-struktur-proyek/cashier/cmd/cashier` |
| Kebiasaan Baik, Test, dan Konkurensi | Gaya kode: gofmt, go vet, penamaan, dan komentar | `.` | `gofmt -l . && go vet ./...` |
| Kebiasaan Baik, Test, dan Konkurensi | Menulis test dengan package testing | `07-test-dan-konkurensi/discount/` | `go test -v ./07-test-dan-konkurensi/discount` |
| Kebiasaan Baik, Test, dan Konkurensi | Menulis test dengan package testing | `06-struktur-proyek/cashier/internal/order/` | `go test -v ./06-struktur-proyek/cashier/internal/order` |
| Kebiasaan Baik, Test, dan Konkurensi | Pengenalan goroutine dan channel | `07-test-dan-konkurensi/concurrency/` | `go run -race ./07-test-dan-konkurensi/concurrency` |
| Kebiasaan Baik, Test, dan Konkurensi | Data race dan sync.Mutex | `07-test-dan-konkurensi/mutex/` | `go run -race ./07-test-dan-konkurensi/mutex -unsafe` |
| Kebiasaan Baik, Test, dan Konkurensi | Data race dan sync.Mutex | `07-test-dan-konkurensi/mutex/` | `go test -race ./07-test-dan-konkurensi/mutex` |
| Kebiasaan Baik, Test, dan Konkurensi | context: batas waktu dan pembatalan | `07-test-dan-konkurensi/context/` | `go run ./07-test-dan-konkurensi/context -timeout 2s` |
| Kebiasaan Baik, Test, dan Konkurensi | context: batas waktu dan pembatalan | `07-test-dan-konkurensi/context/` | `go test -race ./07-test-dan-konkurensi/context` |
| Membangun Aplikasi Pertama | Aplikasi 1: membaca input, file, dan flag | `08-aplikasi/word-count/` | `go run ./08-aplikasi/word-count -top 3 README.md` |
| Membangun Aplikasi Pertama | Aplikasi 2: CLI daftar belanja dengan file JSON | `08-aplikasi/shopping-cli/` | `go run ./08-aplikasi/shopping-cli add coffee 2 18000` |
| Membangun Aplikasi Pertama | Aplikasi 3: API HTTP pertama dengan net/http | `08-aplikasi/notes-api/` | `go run ./08-aplikasi/notes-api` |
| Membangun Aplikasi Pertama | Dari ide ke aplikasi: langkah membuat program sendiri | `08-aplikasi/` | `go test ./08-aplikasi/...` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Mengenal LeetCode dan cara memakainya | `09-latihan-soal/` | `go test ./09-latihan-soal/...` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Membaca soal LeetCode dan menerjemahkannya ke Go | `09-latihan-soal/_template/` | `cp -r 09-latihan-soal/_template latihan/two-sum` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Membaca soal LeetCode dan menerjemahkannya ke Go | `latihan/` | `go test ./latihan/two-sum` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Kompleksitas waktu dan memori (Big-O) tanpa rumus rumit | `09-latihan-soal/02-two-sum/` | `go test -v ./09-latihan-soal/02-two-sum` |
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
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 4: sekali jalan dan pemrograman dinamis sederhana | `09-latihan-soal/01-fizz-buzz/` | `go test ./09-latihan-soal/01-fizz-buzz` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 4: sekali jalan dan pemrograman dinamis sederhana | `09-latihan-soal/10-best-time-to-buy-and-sell-stock/` | `go test ./09-latihan-soal/10-best-time-to-buy-and-sell-stock` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 4: sekali jalan dan pemrograman dinamis sederhana | `09-latihan-soal/12-climbing-stairs/` | `go test -v ./09-latihan-soal/12-climbing-stairs` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 5: linked list dan binary tree | `09-latihan-soal/14-reverse-linked-list/` | `go test ./09-latihan-soal/14-reverse-linked-list` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 5: linked list dan binary tree | `09-latihan-soal/15-merge-two-sorted-lists/` | `go test ./09-latihan-soal/15-merge-two-sorted-lists` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Pola 5: linked list dan binary tree | `09-latihan-soal/16-maximum-depth-of-binary-tree/` | `go test ./09-latihan-soal/16-maximum-depth-of-binary-tree` |
| Menyelesaikan Soal LeetCode Tingkat Easy | Rencana latihan dan cara tidak buntu | `09-latihan-soal/` | `go test ./09-latihan-soal/...` |
| Penutup dan Uji Pengetahuan Akhir | Rangkuman dan langkah berikutnya | `.` | `go test ./...` |

## Perintah yang sering dipakai

```bash
go run ./02-dasar-bahasa/variables  # run one example
go test ./...                       # every test
go test -race ./...                 # with the race detector
gofmt -l .                          # files that are not formatted yet
go vet ./...                        # checks for common mistakes
```
