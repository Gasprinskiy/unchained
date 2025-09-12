FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o main .

FROM alpine:latest
WORKDIR /root/

RUN apk add --no-cache tzdata

ENV TZ=Asia/Tashkent

COPY --from=builder /app/main .

CMD ["./main"]