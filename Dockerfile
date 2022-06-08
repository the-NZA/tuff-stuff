## Step 1
FROM golang:1.18.3-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY ./backend ./

RUN go build -v ./cmd/tuff


## Step 2
FROM scratch

WORKDIR /app

COPY --from=builder /app/tuff ./

CMD ["tuff"]