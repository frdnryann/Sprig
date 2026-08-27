# Database Schema

Dokumentasi skema database untuk aplikasi budgeting (pengelolaan anggaran, kategori, dan pengeluaran/pemasukan pengguna).

## Daftar Isi

- [Overview](#overview)
- [Entity Relationship](#entity-relationship)
- [Tabel](#tabel)
  - [users](#users)
  - [budgets](#budgets)
  - [categories](#categories)
  - [expenses](#expenses)
- [Migrations](#migrations)

## Overview

Database ini terdiri dari 4 tabel utama:

| Tabel | Deskripsi |
|-------|-----------|
| `users` | Data akun pengguna |
| `budgets` | Anggaran bulanan tiap user (metode 50/30/20) |
| `categories` | Kategori transaksi milik masing-masing user |
| `expenses` | Catatan transaksi pengeluaran/pemasukan |

**Engine:** InnoDB
**Charset:** `utf8mb4` — **Collation:** `utf8mb4_unicode_ci`

## Entity Relationship

```
users (1) ──< (N) budgets
users (1) ──< (N) categories
users (1) ──< (N) expenses
categories (1) ──< (N) expenses
```

- Satu user bisa punya banyak budget (satu per kombinasi tahun+bulan).
- Satu user bisa punya banyak kategori, tapi nama kategori harus unik per user.
- Satu user bisa punya banyak expense, dan setiap expense wajib terhubung ke satu kategori.

## Tabel

### `users`

Menyimpan data akun pengguna.

| Kolom | Tipe | Keterangan |
|-------|------|------------|
| `id` | `bigint unsigned` | Primary key, auto increment |
| `name` | `varchar(100)` | Nama pengguna |
| `email` | `varchar(255)` | Wajib unik |
| `password` | `varchar(255)` | Password (harus disimpan dalam bentuk hash) |
| `created_at` | `timestamp` | Default waktu saat baris dibuat |
| `updated_at` | `timestamp` | Otomatis terupdate saat baris diubah |

**Constraint:**
- `UNIQUE (email)`

---

### `budgets`

Menyimpan alokasi anggaran bulanan per user menggunakan metode **50/30/20** (needs/wants/savings).

| Kolom | Tipe | Keterangan |
|-------|------|------------|
| `id` | `bigint unsigned` | Primary key, auto increment |
| `user_id` | `bigint unsigned` | FK ke `users.id` |
| `total_income` | `bigint unsigned` | Total pemasukan pada bulan tersebut |
| `needs_percentage` | `tinyint unsigned` | Default `50` |
| `wants_percentage` | `tinyint unsigned` | Default `30` |
| `savings_percentage` | `tinyint unsigned` | Default `20` |
| `year` | `smallint unsigned` | Tahun anggaran |
| `month` | `tinyint unsigned` | Bulan anggaran (1–12) |
| `created_at` | `timestamp` | Default waktu saat baris dibuat |
| `updated_at` | `timestamp` | Otomatis terupdate saat baris diubah |

**Constraint:**
- `UNIQUE (user_id, year, month)` — satu user hanya boleh punya satu budget per bulan
- `FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE` — budget ikut terhapus jika user dihapus
- `CHECK (month BETWEEN 1 AND 12)`
- `CHECK (needs_percentage + wants_percentage + savings_percentage = 100)` — total alokasi harus 100%

---

### `categories`

Menyimpan kategori transaksi yang dibuat oleh masing-masing user (misalnya: Makan, Transportasi, Gaji, dll).

| Kolom | Tipe | Keterangan |
|-------|------|------------|
| `id` | `bigint unsigned` | Primary key, auto increment |
| `user_id` | `bigint unsigned` | FK ke `users.id` |
| `name` | `varchar(100)` | Nama kategori |
| `created_at` | `timestamp` | Default waktu saat baris dibuat |
| `updated_at` | `timestamp` | Otomatis terupdate saat baris diubah |

**Constraint:**
- `UNIQUE (user_id, name)` — nama kategori tidak boleh duplikat untuk user yang sama (tapi user lain boleh punya nama kategori yang sama)
- `FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE` — kategori ikut terhapus jika user dihapus

---

### `expenses`

Menyimpan catatan transaksi, baik pengeluaran maupun pemasukan.

| Kolom | Tipe | Keterangan |
|-------|------|------------|
| `id` | `bigint unsigned` | Primary key, auto increment |
| `user_id` | `bigint unsigned` | FK ke `users.id` |
| `category_id` | `bigint unsigned` | FK ke `categories.id` |
| `amount` | `bigint unsigned` | Nominal transaksi |
| `description` | `varchar(255)` | Deskripsi transaksi (opsional) |
| `expense_date` | `date` | Tanggal transaksi |
| `type` | `enum('pengeluaran','pemasukan')` | Jenis transaksi |
| `created_at` | `timestamp` | Default waktu saat baris dibuat |
| `updated_at` | `timestamp` | Otomatis terupdate saat baris diubah |

**Index:**
- `idx_expenses_user_date (user_id, expense_date)` — mempercepat query riwayat transaksi per user berdasarkan tanggal
- `idx_expenses_category (category_id)` — mempercepat query filter berdasarkan kategori

**Constraint:**
- `FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE` — expense ikut terhapus jika user dihapus
- `FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT` — kategori **tidak bisa dihapus** selama masih dipakai oleh expense (harus pindahkan/hapus expense-nya dulu)

## Migrations

Migration database dikelola menggunakan [golang-migrate](https://github.com/golang-migrate/migrate) yang dijalankan via Docker Compose, dan dipersingkat lewat `Makefile`.

Pastikan file `.env` berisi variabel berikut sebelum menjalankan perintah migrasi:

```env
DB_USERNAME=root
DB_PASSWORD=secret
DB_HOST=db
DB_PORT=3306
DB_NAME=nama_database
```

### Membuat file migration baru

```bash
sudo make migrate-create name=create_users
```

Perintah ini akan membuat dua file baru di folder `/migrations`: `..._up.sql` dan `..._down.sql`.

### Menjalankan migration (apply)

```bash
make migrate-up
```

Menjalankan seluruh migration yang belum diterapkan ke database.

### Membatalkan migration terakhir (rollback)

```bash
make migrate-down
```

Membatalkan 1 migration paling terakhir yang sudah diterapkan.

### Memaksa versi migration tertentu

```bash
make migrate-force v=<nomor_versi>
```

Digunakan saat migration berada dalam status "dirty" (gagal di tengah jalan) dan perlu direset paksa ke versi tertentu.

### Mengecek versi migration saat ini

```bash
make migrate-version
```

Menampilkan versi migration terakhir yang sudah diterapkan ke database.