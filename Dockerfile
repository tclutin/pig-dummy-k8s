FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main main.go

FROM scratch

COPY --from=builder /app/main /main
COPY --from=builder /app/resources /resources

CMD ["/main"]