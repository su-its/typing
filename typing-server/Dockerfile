# 基本イメージ
FROM golang:1.22.2 as builder

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# 依存関係をインストール
RUN go mod download

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ./api/cmd/main.go

# 実行イメージ
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# tzdataパッケージのインストール
RUN apk --no-cache add tzdata

# curlのインストール
RUN apk --no-cache add curl

WORKDIR /root

# ビルドしたバイナリをコピー
COPY --from=builder /app/server .
COPY --from=builder /app/.env .


# アプリケーションの実行
CMD ["./server","-seed","false"]
