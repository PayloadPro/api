# Provision the app
FROM golang:1.11-rc-alpine
WORKDIR /go/src/github.com/andrew-waters/payload.pro
COPY . .
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o payload.pro .

# Build the deployable container
FROM golang:1.11-rc-alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/src/github.com/andrew-waters/payload.pro .
CMD ["./payload.pro"]  
