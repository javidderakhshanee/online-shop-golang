FROM  golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY *.yaml ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .


CMD [ "./main" ]
