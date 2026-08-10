# Website Sekumlah

Tanpa basa basi, website ini digunakan buat apa (Udah jelas dari judulnya juga). Langsung ke bagian teknisnya saja


## Struktur Project

```
root-folder
-> cmd
    -> main.go // main file (tempat semua nya dirakit)
-> internal
    -> announcements // bagian dari fitur website (CRUD)
    -> config // konfigurasi dasar yang diambil dari .env
    -> model // model database
    -> router // yaa.., ini router
    -> teachers // bagian dari fitur website (CRUD)
-> migration
-> static
-> views
-> go.mod
-> go.sum
-> tailwind.config.js // standaalone CLI installation
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
```
git clone https://github.com/frdnryann/website-sekumlah.git
```

#### 2. Install Tailwindcss (V3.4.19)
```
// Download file binary tailwindcss (Sesuaikan dengan OS yang dipakai)
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.19/tailwindcss-linux-x64

// izinkan pengeksekusian file
chmod +x tailwindcss-linux-x64

// ubah nama
mv tailwindcss-linux-x64 tailwindcss
```