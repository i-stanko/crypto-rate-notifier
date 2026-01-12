# Crypto Rate Notifier

Crypto Rate Notifier is a lightweight HTTP service written in Go that provides the current Bitcoin exchange rate and allows users to subscribe to rate notifications via email. The project is designed as a clean, container-ready service with a clear separation of concerns and production-oriented structure.

[![asciicast](https://asciinema.org/a/767442.svg)](https://asciinema.org/a/767442)


## Features

- Fetches current Bitcoin rate from an external API ([CoinGecko](https://www.coingecko.com/))
- Email subscription management
- Simple REST API
- File-based storage (easy to replace with DB)
- Docker-ready (multi-stage build)
- Clean Go project structure


## Requirements

- Go 1.23+
- Docker (optional)


## Getting Started

- **Clone repository**
```bash
git clone https://github.com/i-stanko/crypto-rate-notifier.git
cd crypto-rate-notifier
```

- **Run locally**
```bash
go run ./cmd/server
```

The service will be available at:
- http://localhost:8080



## API Endpoints

- **Get current BTC to UAH exchange rate**
```bash
curl http://localhost:8080/api/rate
```

- **Subscribe email**
```bash
curl -X POST \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "email=test@test.com" \
  http://localhost:8080/api/subscribe
```

- **List subscribed emails**
```bash
curl http://localhost:8080/api/subscribers
```



## Docker

- **Build image**
```bash
docker build -t crypto-rate-notifier .
```

- **Run container**
```bash
docker run -p 8080:8080 crypto-rate-notifier
```

## Notes

- Subscribers are stored in a local file using a storage interface.
- The project is intended as a demo service for Go backend and DevOps practices.
- The codebase is ready for further extension (Tests, Kubernetes, Helm, CI/CD).
