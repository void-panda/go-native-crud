## Go Native CRUD - Project Pembelajaran

Aplikasi ini merupakan project sederhana untuk pembelajaran implementasi operasi CRUD (Create, Read, Update, Delete) menggunakan bahasa pemrograman Go (Golang) **tanpa menggunakan framework**.

### Teknologi
- Golang v.1.24.3
- MySQL 8
- Bootstrap 5

### Fitur Utama
- Implementasi CRUD dasar pada dua tabel yang saling berelasi: **Post** dan **Category**.
- Setiap Post memiliki relasi ke satu Category.
- Operasi CRUD meliputi: tambah, lihat, ubah, dan hapus data pada kedua tabel.
- Tidak terdapat sistem autentikasi, sehingga aplikasi ini fokus pada logika dasar pengelolaan data.

### Tujuan Project
Project ini dibuat sebagai sarana belajar untuk memahami:
- Cara membangun aplikasi CRUD native di Go tanpa framework eksternal.
- Cara mengelola relasi antar tabel di database menggunakan Go.
- Struktur kode dan best practice sederhana dalam pengembangan aplikasi Go.

### Struktur Database (Contoh)

**Category**
- id (int, primary key)
- name (string)

**Post**
- id (int, primary key)
- title (string)
- content (text)
- category_id (int, foreign key ke Category)

### Catatan
- Project ini **tidak** memiliki sistem autentikasi.
- Cocok untuk pemula yang ingin memahami dasar CRUD dan relasi tabel di Go.

---
Selamat belajar dan bereksperimen dengan Go Native CRUD!
