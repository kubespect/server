FROM golang AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o server main.go

FROM debian

WORKDIR /app

COPY --from=builder /app/server /app/server

ARG WS_PORT=8080

ARG GRPC_PORT=9090

ENV WS_PORT=$WS_PORT

ENV GRPC_PORT=$GRPC_PORT

EXPOSE $WS_PORT
EXPOSE $GRPC_PORT

CMD ["./server"]