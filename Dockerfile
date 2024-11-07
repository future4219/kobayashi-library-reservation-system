# ベースイメージとしてGoを指定
FROM golang:1.21.3 AS builder

# 作業ディレクトリを設定
WORKDIR /app

# モジュールファイルをコピーして、依存関係をインストール
COPY go.mod go.sum ./
RUN go mod tidy

# ソースコードをコンテナにコピー
COPY ./api .

# Goアプリケーションをビルド
RUN go build -o .

# 実行用の軽量なイメージに切り替え
FROM debian:bullseye-slim

# 作業ディレクトリを設定
WORKDIR /root/

# ビルドしたアプリケーションをコピー
COPY --from=builder /app/main .

# ポート8080を開放
EXPOSE 8080

# コンテナ起動時に実行するコマンドを指定
CMD ["./main"]