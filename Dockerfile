FROM golang:1.14.0-alpine AS builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh ca-certificates

# Install go dep
RUN wget https://raw.githubusercontent.com/golang/dep/master/install.sh
RUN chmod +x ./install.sh
RUN ["./install.sh"]

WORKDIR /go/src/TSACodingChallengeAPI/src
COPY src/Gopkg.lock src/Gopkg.toml ./
RUN dep ensure --vendor-only

COPY src/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app -ldflags -s
RUN CGO_ENABLED=0 go test -v -short ./...

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app  /app

EXPOSE 10010

ENTRYPOINT ["./app"]