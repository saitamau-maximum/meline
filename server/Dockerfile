
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main .


FROM alpine AS prod
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/sql /app/sql

EXPOSE 8000
ENTRYPOINT ["/app/main"]


FROM golang:1.20-alpine AS dev

ENV CGO_ENABLED 0
WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8000

CMD ["go", "run", "main.go"]
