# 基本イメージ
FROM golang:1.23.4 AS builder

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# 依存関係をインストール
RUN go mod download

RUN cd internal/infra/ent && go generate

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ./cmd/server/main.go

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

# アプリケーションの実行
CMD ["./server"]
