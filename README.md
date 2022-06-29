# forum
Forum with clean architecture

## How to run?
```bash
go run ./cmd/
```

## Run on Docker
1. Build image
```bash
docker build . -t forum-image
```
2. Run container
```bash
docker run -p 80:8080 --rm --name forum-container forum-image
```