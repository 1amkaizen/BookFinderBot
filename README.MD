# BookFinderBot

BookFinderBot adalah sebuh bot Telegram yang membantu pengguna mencari dan menemukan buku dan ebook berdasarkan judul atau topik yang diinginkan. Bot ini juga menyediakan fitur untuk mendapatkan link ulasan buku yang diinginkan.

## Cara Menggunakan Bot

### Memperoleh Token Bot Telegram

1. Buka aplikasi Telegram.
2. Cari bot bernama "BotFather".
3. Mulai percakapan dengan BotFather.
4. Ketik perintah `/newbot` untuk membuat bot baru.
5. Ikuti instruksi dari BotFather untuk memberi nama dan username baru untuk bot Anda.
6. BotFather akan memberikan token bot Anda. Simpan token ini dengan aman.

### Export Token Bot Telegram

1. Buka terminal atau command prompt.
2. Export token bot Anda sebagai environment variable dengan menjalankan perintah berikut:
    ```bash
    export TELEGRAM_BOT_TOKEN="YOUR_BOT_TOKEN"
    ```
   Ganti `YOUR_BOT_TOKEN` dengan token bot yang Anda peroleh dari BotFather.

### Menyiapkan Data Produk dan Link Ulasan

1. Ganti isi file `products.txt` dengan produk-produk yang ingin Anda tampilkan dalam bot. Format setiap baris adalah `Nama Produk: https://linkproduk`.
2. Ganti isi file `link_reviews.txt` dengan link ulasan untuk setiap produk. Format setiap baris adalah `Nama Produk: https://linkulasan`.
3. Pastikan nama produk di `link_reviews.txt` cocok dengan nama produk di `products.txt`.

### Menjalankan Bot

Pastikan Anda memiliki Go (Golang) terinstal di komputer Anda sebelum menjalankan bot.

1. Clone repositori ini ke komputer Anda.
2. Buka terminal atau command prompt.
3. Arahkan ke direktori repositori bot Anda.
4. Jalankan perintah `go run main.go` untuk menjalankan bot.

## Perintah yang Tersedia

- `/start`: Memulai percakapan dengan bot dan menampilkan pesan selamat datang.
- `/help`: Menampilkan daftar perintah yang tersedia beserta contoh penggunaannya.
- `/ulasan [nama lengkap produk]`: Mendapatkan link ulasan produk yang diinginkan.

## Mengkontribusi

Anda dapat berkontribusi pada pengembangan CarieBookBot dengan melakukan pull request ke repositori ini. Silakan buka issue untuk saran atau permintaan fitur.

## Mengirim Ulasan

Anda juga dapat memberikan ulasan langsung melalui [form ulasan kami](https://aigoretech.rf.gd/kirim-ulasan).

## Lisensi

[MIT](LICENSE)
