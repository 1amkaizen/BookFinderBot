# Gunakan base image untuk Go
FROM golang:1.19-alpine

# Set environment variables
ENV GO111MODULE=on

# Buat directory kerja
WORKDIR /app

# Copy semua file ke directory kerja
COPY . .

# Unduh dependencies
RUN go mod tidy

# Build aplikasi Go
RUN go build -o BookFinderBot .

# Jalankan aplikasi
CMD ["./BookFinderBot"]
