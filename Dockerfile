# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.19.0-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum /
RUN go mod download

COPY docker .

RUN go build -trimpath -ldflags "-w -s" -o app

# --------------------------------------------

# デプロイ用のコンテナ

From debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# --------------------------------------------

# ローカル環境で利用するホットリロード環境

FROM golang:1.19.0 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]