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


## 📤 Upload Gambar Produk

- **Endpoint:** `POST /upload/:product_id`
- **Deskripsi:** Mengunggah gambar produk berdasarkan ID.
- **Request:**
  - Form Data:
    - Key: `image` (type: File)
- **Validasi:**
  - Format: `.jpg`, `.jpeg`, `.png`
  - Maks ukuran: 2MB
- **Response (200):**

```json
{
  "message": "File berhasil diupload",
  "filename": "1.jpg",
  "image_url": "/uploads/1.jpg"
}

## 📤 Download Gambar Produk

- **Endpoint:** `GET /download/:product_id`
- **Deskripsi:** Mengunduh gambar produk berdasarkan ID.
- **Response:** File gambar

```json
{
  "message": "File berhasil diupload",
  "filename": "1.jpg",
  "image_url": "/uploads/1.jpg"
}