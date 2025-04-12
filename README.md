# inventory-api

API backend untuk manajemen inventaris yang dibangun menggunakan Golang dengan framework Gin dan GORM untuk integrasi dengan MySQL. API ini menyediakan endpoint untuk mengelola produk, inventaris, dan pesanan.

## Setup Database

1. **Instalasi MySQL:**
   - Pastikan MySQL sudah terinstal di komputer kamu. Kamu bisa mengunduh MySQL dari [situs resmi MySQL](https://dev.mysql.com/downloads/installer/).
   
2. **Setup Database**
   - Buat database
   - Sesuaikan konfigurasi database kamu di file `.env`
  
3. **Menjalankan Project**
   - Pastikan kamu sudah install semua dependecy yang digunakan di proyek. bisa menggunakan perintah "go mod tidy"
   - Jalankan aplikasi menggunakan `go run main.go`
   - API akan berjalan di http://localhost:8080

4. **DOKUMENTASI API**
   - Untuk dokumentasi lebih lanjut, kamu bisa mengakses dokumentasi API secara online di (https://documenter.getpostman.com/view/40938878/2sB2cYeLxw)
