# Gunakan gambar dasar resmi Golang
FROM golang:1.20 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Buat direktori kerja
WORKDIR /BookFinderBot

# Copy kode sumber dari proyek di GitHub ke direktori kerja dalam container
COPY . .

# Unduh dan instal dependensi, lalu build aplikasi
RUN go mod tidy
RUN go build -o /BookFinderBot .

# Gunakan gambar dasar yang lebih kecil untuk hasil akhir
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Copy executable dari tahap builder
COPY --from=builder /BookFinderBot /BookFinderBot

# Set environment variable untuk token bot Telegram
ENV TELEGRAM_BOT_TOKEN="7342847814:AAEIkAumiS5nm0Eq5yldKsbXaL_nGSAWBK4" 

# Jalankan bot saat container dimulai
CMD ["/BookFinderBot"]
