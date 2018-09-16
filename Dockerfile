FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM golang:1.11-rc-alpine as build
WORKDIR /go/src/github.com/PayloadPro/pro.payload.api
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY ./ /go/src/github.com/PayloadPro/pro.payload.api
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o payload.pro .

FROM scratch
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/github.com/PayloadPro/pro.payload.api/payload.pro /bin/payload.pro
CMD ["./bin/payload.pro"]