# ---------- build stage ----------
FROM golang:1.23-alpine AS builder

WORKDIR /app

# install git (needed for go modules sometimes)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server

# ---------- runtime stage ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/server /app/server
COPY subscribers.txt /app/subscribers.txt

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
