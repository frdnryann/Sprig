# Go-Financial (Nama sementara)

Aplikasi untuk mengatur keuangan yang sebelumnya tertunda proses pembuatannya. Kini dibuat ulang dengan menggunakan bahasa pemrograman Golang

# Features

Stay Tuned ya.., heheheha


## Struktur Project

```
root-folder
-> cmd
    -> web
        -> main.go // main file (tempat semua nya dirakit)
-> docker
    -> golang
        -> Dockerfile
    -> mysql
        -> my.cnf
-> internal
    -> config // konfigurasi dasar yang diambil dari .env
    -> features // bagian dari fitur website (CRUD)
        -> feat-1
        -> feat-2
    -> model // model database
    -> router // yaa.., ini router
-> migration
-> static
-> views
-> go.mod
-> go.sum
-> tailwind.config.js // standalone CLI installation
-> ...
```

## Tools yang digunakan

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
        <td><a href="https://tailwindcss.com/blog/standalone-cli">Tailwindcss (Standalone CLI installation)</a></td>
        <td><a href="https://docs.docker.com/desktop/setup/install/windows-install/">Docker</a></td>
    </tr>
    <tr>
        <td><a href="https://quilljs.com/docs/quickstart">Quill</a></td>
        <td><a href="https://github.com/golang-migrate/migrate">Golang-migrate</a></td>
    </tr>
</table>

## Instalasi

Pastikan dikomputer mu sudah terinstall wsl, dan juga distro linux pilihanmu. Jika belum terinstall, silahkan install terlebih dahulu dengan membaca dokumentasi resmi dari microsoft [WSL Install](https://learn.microsoft.com/en-us/windows/wsl/install)

##### 1. Clone Repository ini
```bash
git clone https://github.com/frdnryann/Go-financial.git
```

#### 2. Install Image & Container docker untuk project
```bash
sudo docker compose up -d --build

# Cara menghentikan container yang aktif
sudo docker compose stop

# Ingin sekalian menghapus container ?
sudo docker compose down
``` 

##### 3. Migrate Database
```bash
# Install dependensi yang dibutuhkan
sudo apt install make

# Migrate database
sudo make migrate-up
```
> NOTE : Untuk detail penggunaan command migrasi database, silahkan baca list yang ada di dokumentasi ini [Migration Command List](/migrations/README.md)

#### 4. Install package yang dibutuhkan golang
```bash
go mod download

# Ingin package nya disimpan di local ?
go mod vendor
```

#### 5. Install Tailwindcss (V3.4.19)
```bash
# Download file binary tailwindcss (Sesuaikan dengan OS yang dipakai)
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.19/tailwindcss-linux-x64

# izinkan pengeksekusian file
chmod +x tailwindcss-linux-x64

# ubah nama
mv tailwindcss-linux-x64 tailwindcss
```