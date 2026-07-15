# Portal News Blog API

Backend API untuk aplikasi portal berita/blog yang dibangun dengan Go. Project ini disiapkan menggunakan pendekatan Hexagonal Architecture atau Ports and Adapters agar business logic tidak bergantung langsung pada detail eksternal seperti HTTP handler, database, storage, atau service pihak ketiga.

> Status: project masih tahap awal. Struktur dasar sudah ada, tetapi entry point aplikasi, koneksi database, migration, handler, repository, domain, dan service belum lengkap.

## Tujuan Arsitektur

Hexagonal Architecture membantu project tetap mudah dirawat ketika aplikasi mulai berkembang. Inti aplikasi diletakkan di domain dan service, sedangkan teknologi luar seperti database, HTTP framework, JWT, atau Cloudflare dibuat sebagai adapter.

Dengan pola ini, perubahan detail teknis seperti mengganti database driver, mengganti HTTP router, atau menambah storage provider tidak perlu mengubah business logic utama.

## Struktur Project

```text
portal-news-blog/
|-- cmd/
|   `-- main.go              # Entry point aplikasi. Belum tersedia.
|-- config/
|   `-- config.go            # Loader konfigurasi dari environment. Masih kosong.
|-- database/
|   `-- migrations/          # File migration untuk go migrate. Belum tersedia.
|-- internal/
|   |-- adapter/
|   |   |-- handler/         # HTTP handler/controller.
|   |   |-- repository/      # Query dan akses database.
|   |   `-- cloudflare/      # Integrasi Cloudflare atau storage eksternal.
|   |-- app/
|   |   `-- app.go           # Inisialisasi dependency utama aplikasi.
|   `-- core/
|       |-- domain/          # Entity dan model bisnis utama.
|       `-- service/         # Business logic aplikasi.
|-- lib/
|   |-- conf/                # Helper konfigurasi dan konversi nilai.
|   `-- jwt/                 # Generate dan verifikasi JWT.
|-- .env                     # Konfigurasi lokal. Jangan commit secret.
|-- go.mod
`-- README.md
```

Beberapa folder pada struktur di atas masih berupa rencana pengembangan. Saat ini repository baru berisi `config/`, `internal/app/`, `lib/conf/`, `go.mod`, dan `README.md`.

## Alur Layer

```text
HTTP Request
    |
    v
adapter/handler
    |
    v
core/service
    |
    v
core/domain
    |
    v
adapter/repository
    |
    v
Database
```

Handler bertugas menerima request dan mengubahnya menjadi input service. Service menjalankan aturan bisnis. Repository hanya fokus pada akses data. Domain menyimpan struktur data inti seperti user, artikel, kategori, tag, atau komentar.

## Prasyarat

- Go `1.22.2` atau versi kompatibel.
- PostgreSQL atau database lain sesuai implementasi yang nanti dipilih.
- `golang-migrate` jika migration sudah tersedia.

## Environment

Buat file `.env` di root project. Contoh variabel yang saat ini digunakan atau direncanakan:

```env
APP_ENV=development
APP_PORT=8080

DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=password
DATABASE_NAME=portal_news_blog
DATABASE_MAX_OPEN_CONNECTIONS=10
DATABASE_MAX_IDLE_CONNECTIONS=5

JWT_SECRET_KEY=change-this-secret
JWT_ISSUER=portal-news-blog
```

Catatan: file `.env` saat ini memakai nama `DATABSE_NAME`. Sebaiknya ubah menjadi `DATABASE_NAME` saat loader konfigurasi dibuat agar konsisten dengan variabel database lainnya.

## Setup Lokal

Clone repository dan masuk ke folder project:

```bash
git clone <repository-url>
cd portal-news-blog
```

Install dependency Go:

```bash
go mod tidy
```

Siapkan `.env`:

```bash
cp .env.example .env
```

Jika `.env.example` belum tersedia, buat `.env` manual berdasarkan contoh pada bagian Environment.

## Menjalankan Aplikasi

Entry point aplikasi belum tersedia. Setelah `cmd/main.go` dibuat, aplikasi dapat dijalankan dengan pola berikut:

```bash
go run ./cmd
```

Atau jika file entry point dibuat langsung sebagai `cmd/main.go`:

```bash
go run ./cmd/main.go
```

## Migration Database

Folder migration belum tersedia. Setelah migration dibuat, perintah yang disarankan:

```bash
migrate create -ext sql -dir database/migrations -seq create_users_table
```

Menjalankan migration:

```bash
migrate -path database/migrations \
  -database "postgres://user:password@localhost:5432/portal_news_blog?sslmode=disable" \
  up
```

Rollback migration terakhir:

```bash
migrate -path database/migrations \
  -database "postgres://user:password@localhost:5432/portal_news_blog?sslmode=disable" \
  down 1
```

## Rekomendasi Pengembangan Berikutnya

1. Buat loader konfigurasi di `config/config.go`.
2. Tambahkan entry point aplikasi di `cmd/main.go`.
3. Tentukan HTTP router yang akan digunakan.
4. Buat koneksi database dan migration awal.
5. Definisikan domain utama seperti `User`, `Article`, `Category`, `Tag`, dan `Comment`.
6. Tambahkan repository dan service untuk fitur autentikasi.
7. Tambahkan JWT helper di `lib/jwt`.

## Konvensi

- Simpan business logic di `internal/core/service`.
- Simpan entity utama di `internal/core/domain`.
- Simpan akses database di `internal/adapter/repository`.
- Simpan handler HTTP di `internal/adapter/handler`.
- Hindari membaca `.env` langsung dari banyak tempat. Gunakan satu loader konfigurasi.
- Jangan commit credential asli, token, atau secret dari `.env`.

## License

Belum ditentukan.

Notes:

- JWT adalah signed token yang membawa informasi (claims) pengguna untuk proses authentication dan authorization tanpa perlu menyimpan session di server.

- Viper adalah library konfigurasi (configuration management) yang sangat populer di Golang. Tujuannya adalah memudahkan aplikasi membaca konfigurasi dari berbagai sumber tanpa Anda perlu menulis kode parsing sendiri.
