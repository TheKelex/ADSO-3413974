FROM golang:1.22-alpine AS build

WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN go build -o /out/horarios-api ./cmd/server

FROM alpine:3.20

RUN adduser -D -H appuser
WORKDIR /app
COPY --from=build /out/horarios-api /app/horarios-api

USER appuser
EXPOSE 8080
CMD ["/app/horarios-api"]

