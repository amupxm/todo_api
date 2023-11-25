FROM golang:1.21.4-alpine3.17 as builder
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app

FROM alpine
# COPY --from=builder /go/bin/app/devConfig.env /go/bin/app/devConfig.env
COPY --from=builder /go/bin/app /go/bin/app
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]
