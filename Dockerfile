## Step 1
FROM golang:1.18.3 as builder

ADD backend /app/backend

WORKDIR /app

COPY go.* ./

RUN go mod download

RUN go build -v -o tuff ./backend/cmd/tuff

EXPOSE 8080

ENTRYPOINT ["./tuff"]


## Step 2
#FROM alpine:3.16 as runner
#
#WORKDIR /app
#
#COPY --from=builder /app/tuff ./tuff
#
#ENTRYPOINT ["./tuff"]