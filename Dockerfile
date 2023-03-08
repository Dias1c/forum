# BUILDING PROJECT
FROM golang:1.18.3 AS builder
LABEL container for staging build container
WORKDIR /src
COPY . .
RUN GOOS=linux go build -o main.exe ./cmd/

# DOCKER STAGE: COPY NEEDED ELEMENTS TO NEW CONTAINER
FROM ubuntu:20.04
LABEL run forum builded file on new container
WORKDIR /src

COPY --from=builder /src/database database
COPY --from=builder /src/web web
COPY --from=builder /src/*.db .
COPY --from=builder /src/configs.env .
COPY --from=builder /src/main.exe .
EXPOSE 8080

CMD ["./main.exe"]