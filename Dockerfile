# 构建阶段
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download

COPY . .

RUN apk add build-base

RUN CGO_ENABLED=1 go build -o vega-server ./cmd/server

# 运行阶段
FROM alpine

WORKDIR /app

COPY --from=builder /app/vega-server .

COPY --from=builder /app/config ./config

COPY --from=builder /app/storage ./storage

EXPOSE 8080

CMD ["./vega-server", "-path=config/dev.yaml"]