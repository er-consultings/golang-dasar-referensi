// Package fizzbuzz: untuk setiap angka 1 sampai n, tulis "Fizz" bila habis dibagi 3,
// "Buzz" bila habis dibagi 5, "FizzBuzz" bila habis dibagi keduanya, selain itu angkanya.
//
// Mirip LeetCode 412. Pola: perulangan dan percabangan.
// Waktu O(n), ruang O(n) untuk hasilnya.
package fizzbuzz

import "strconv"

func fizzBuzz(n int) []string {
	out := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0: // periksa yang paling khusus lebih dulu
			out = append(out, "FizzBuzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}
