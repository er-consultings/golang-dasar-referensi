# Latihan soal

Setiap folder berisi satu soal tingkat Easy yang mirip soal LeetCode, ditulis ulang dengan bahasa sendiri.

- `problem.go` berisi deskripsi soal, pola yang dipakai, kompleksitas, dan solusi.
- `problem_test.go` berisi contoh dari soal dalam bentuk *table-driven test*.

| Folder | Soal | No. LeetCode | Pola |
|---|---|---|---|
| `01-fizz-buzz` | Fizz Buzz | 412 | *loop* dan percabangan |
| `02-two-sum` | Two Sum | 1 | *hash map* |
| `03-contains-duplicate` | Contains Duplicate | 217 | *hash set* |
| `04-valid-anagram` | Valid Anagram | 242 | menghitung frekuensi |
| `05-valid-palindrome` | Valid Palindrome | 125 | *two pointers* |
| `06-reverse-string` | Reverse String | 344 | *two pointers*, *in place* |
| `07-merge-sorted-array` | Merge Sorted Array | 88 | *two pointers* dari belakang |
| `08-valid-parentheses` | Valid Parentheses | 20 | *stack* |
| `09-binary-search` | Binary Search | 704 | *binary search* |
| `10-best-time-to-buy-and-sell-stock` | Best Time to Buy and Sell Stock | 121 | *single pass*, simpan nilai terkecil |
| `11-move-zeroes` | Move Zeroes | 283 | *two pointers*: baca dan tulis |
| `12-climbing-stairs` | Climbing Stairs | 70 | rekursi ke *dynamic programming* |
| `13-roman-to-integer` | Roman to Integer | 13 | *map* dan melihat karakter berikutnya |
| `14-reverse-linked-list` | Reverse Linked List | 206 | *linked list*, tiga *pointer* |
| `15-merge-two-sorted-lists` | Merge Two Sorted Lists | 21 | *linked list*, *dummy node* |
| `16-maximum-depth-of-binary-tree` | Maximum Depth of Binary Tree | 104 | rekursi pada *tree* |

## Cara berlatih

1. Baca deskripsi di `problem.go`, **jangan lihat solusinya dulu**.
2. Salin *template* ke folder latihan Anda sendiri:
   ```bash
   cp -r 09-latihan-soal/_template latihan/two-sum
   ```
3. Tulis contoh soal sebagai *test* lebih dulu, lalu solusinya.
4. Jalankan `go test ./latihan/two-sum`.
5. Setelah lulus, bandingkan dengan solusi di folder ini dan perhatikan kompleksitasnya.
6. *Submit* jawaban ke LeetCode. Di sana cukup salin fungsinya saja, tanpa `package` dan `import` yang tidak dipakai.

```bash
go test ./09-latihan-soal/...          # run everything
go test -v ./09-latihan-soal/02-two-sum
```
