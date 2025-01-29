# Dockerfile
# Start by using the official Go image as the base image
FROM golang:1.23.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# # Copy the Go module files
COPY go.mod ./

# # Download the module dependencies
RUN go mod download

# air を正しいパスでインストール
RUN go install github.com/air-verse/air@latest

# # Copy the source code
COPY . ./

# デフォルトのコマンドを設定
# コンテナ起動後は無限にスリープし、コンテナを終了させない（開発時の利便性向上）
CMD ["sleep", "infinity"]

