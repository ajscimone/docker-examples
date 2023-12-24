FROM golang:latest AS builder

WORKDIR /go/src/

COPY . .

RUN go build -o /go/src/app .

FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /src/
COPY --from=builder /go/src/app /src/app

CMD ["/src/app"]
