## Step 1
FROM golang:1.18.3 as builder

ADD backend /app/backend

WORKDIR /app

COPY go.* ./

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -v -o tuff -a -ldflags '-linkmode external -extldflags "-static"' ./backend/cmd/tuff

## Step 2
FROM scratch

WORKDIR /app

COPY --from=builder /app/tuff .

EXPOSE 8080

ENTRYPOINT ["./tuff"]