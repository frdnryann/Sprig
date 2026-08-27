# Go-Financial (Nama sementara)

Aplikasi pengelolaan keuangan yang sempat tertunda pengerjaannya, kini dibangun ulang menggunakan bahasa pemrograman Golang.

## Daftar Isi

- [Features](#features)
- [Struktur Project](#struktur-project)
- [Tools yang Digunakan](#tools-yang-digunakan)
- [Instalasi](#instalasi)
  - [1. Clone Repository](#1-clone-repository-ini)
  - [2. Install Image & Jalankan Container Docker](#2-install-image--jalankan-container-docker-untuk-project)
  - [3. Migrasi Database](#3-migrasi-database)
  - [4. Install Package yang Dibutuhkan](#4-install-package-yang-dibutuhkan-golang)
  - [Tambahan: Go Unit Test](#tambahan--go-unit-test)

## Features

- **Income Management** — Mencatat dan mengelola pemasukan.
- **Expense Management** — Mencatat pengeluaran berdasarkan transaksi.
- **Category Management** — Mengelompokkan transaksi berdasarkan kategori.
- **Financial Summary** — Melihat ringkasan kondisi keuangan.
- **Transaction Search & Filtering** — Mencari dan memfilter transaksi berdasarkan kriteria tertentu.
- **Transaction History** — Melihat riwayat pemasukan dan pengeluaran.
- **Automatic Balance Calculation** — Menghitung saldo secara otomatis berdasarkan transaksi.
- **Budget Management** — Menentukan dan memantau batas pengeluaran.
- **Financial Statistics** — Menampilkan statistik pemasukan dan pengeluaran.
- **User Authentication** — Mengamankan akses data setiap pengguna.
- **User Management** — Mengelola informasi akun pengguna.

## Struktur Project

```
root-folder
├── cmd
│   └── web
│       └── main.go        // entry point (tempat semua komponen dirakit)
├── docker
│   ├── golang
│   │   └── Dockerfile
│   └── mysql
│       └── my.cnf
├── internal
│   ├── config              // konfigurasi dasar yang diambil dari .env
│   ├── features             // masing-masing fitur aplikasi (CRUD)
│   │   ├── feat-1
│   │   └── feat-2
│   ├── model                // model database
│   └── router                // definisi routing
├── migration
├── go.mod
├── go.sum
└── ...
```

## Tools yang Digunakan

<table>
    <tr>
        <th>Frontend</th>
        <th>Backend</th>
    </tr>
    <tr>
        <td><a href="https://htmx.org/docs/">HTMX</a></td>
        <td><a href="https://go.dev/">Golang (Vanilla)</a></td>
    </tr>
    <tr>
        <td><a href="https://templ.guide/">Templ</a></td>
        <td><a href="https://github.com/go-sql-driver/MYSQL">MySQL</a></td>
    </tr>
    <tr>
        <td><a href="https://tailwindcss.com/blog/standalone-cli">Tailwind CSS (Standalone CLI installation)</a></td>
        <td><a href="https://docs.docker.com/desktop/setup/install/windows-install/">Docker</a></td>
    </tr>
    <tr>
        <td><a href="https://quilljs.com/docs/quickstart">Quill</a></td>
        <td><a href="https://github.com/golang-migrate/migrate">Golang-migrate</a></td>
    </tr>
</table>

## Instalasi

Pastikan di komputer kamu sudah terinstall WSL beserta distro Linux pilihanmu. Jika belum, silakan install terlebih dahulu dengan mengikuti dokumentasi resmi Microsoft: [WSL Install](https://learn.microsoft.com/en-us/windows/wsl/install).

##### 1. Clone Repository ini
```bash
git clone https://github.com/frdnryann/Go-financial.git
```

##### 2. Install Image & Jalankan Container Docker untuk Project
```bash
sudo docker compose up -d --build

# Menghentikan container yang sedang aktif
sudo docker compose stop

# Ingin sekaligus menghapus container?
sudo docker compose down
```

##### 3. Migrasi Database
```bash
# Install dependensi yang dibutuhkan (jika belum terinstall)
sudo apt install make

# Jalankan migrasi database
sudo make migrate-up
```
> **NOTE:** Untuk detail penggunaan command migrasi database lainnya, silakan baca daftar lengkapnya di [Migration Command List](/migrations/README.md).

##### 4. Install Package yang Dibutuhkan Golang
```bash
go mod download

# Ingin package-nya disimpan secara lokal di dalam project?
go mod vendor
```

##### Tambahan : Go Unit Test
```bash
sudo make test path=./file/destination

# Contoh
sudo make test path=./internal/config
```