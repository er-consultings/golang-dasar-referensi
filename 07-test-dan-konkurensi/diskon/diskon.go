// Package diskon adalah contoh kecil untuk belajar table-driven test.
//
//	go test -v ./07-test-dan-konkurensi/diskon
package diskon

// Hitung mengembalikan total setelah dipotong persen.
// Persen nol atau negatif diabaikan, persen di atas 100 dibatasi menjadi 100.
func Hitung(total, persen int64) int64 {
	if persen <= 0 {
		return total
	}
	if persen > 100 {
		persen = 100
	}
	return total - total*persen/100
}
