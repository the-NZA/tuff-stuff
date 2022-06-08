## Step 1
FROM golang:1.18.3 as builder

ADD backend /app/backend

WORKDIR /app

COPY go.* ./

RUN go mod download

#RUN go get -u ./...

RUN go build -v -o tuff ./backend/cmd/tuff


## Step 2
FROM scratch

WORKDIR /app

COPY --from=builder /app/tuff ./

CMD ["tuff"]