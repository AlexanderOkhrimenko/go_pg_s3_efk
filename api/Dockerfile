FROM golang:latest AS builder

RUN mkdir /go/src/api/
WORKDIR /go/src/api/

COPY . .
RUN go mod download
WORKDIR /go/src/api/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/api/cmd/ .

EXPOSE 8080
CMD ["./api"]