# Go-Financial (Nama sementara)

Aplikasi untuk mengatur keuangan yang sebelumnya tertunda proses pembuatannya. Kini dibuat ulang dengan menggunakan bahasa pemrograman Golang

# Featured

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
        <td><a href="#">HTMX</a></td>
        <td><a href="#">Golang (Vanilla)</a></td>
    </tr>
    <tr>
        <td><a href="#">Templ</a></td>
        <td><a href="#">MySQL</a></td>
    </tr>
    <tr>
        <td><a href="#">Tailwindcss (Standalone CLI installation)</a></td>
        <td><a href="#">Docker</a></td>
    </tr>
    <tr>
        <td><a href="#">Quill</a></td>
        <td><a href="#">-</a></td>
    </tr>
</table>

## Instalasi
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

#### 3. Install package yang dibutuhkan golang
```bash
go mod download

# Ingin package nya disimpan di local ?
go mod vendor
```

#### 4. Install Tailwindcss (V3.4.19)
```bash
# Download file binary tailwindcss (Sesuaikan dengan OS yang dipakai)
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.19/tailwindcss-linux-x64

# izinkan pengeksekusian file
chmod +x tailwindcss-linux-x64

# ubah nama
mv tailwindcss-linux-x64 tailwindcss
```