FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/api .

FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
COPY --from=builder /bin/api /bin/api
USER app
EXPOSE 8080
ENTRYPOINT ["/bin/api"]