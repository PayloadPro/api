[![Maintainability](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/maintainability)](https://codeclimate.com/github/codeclimate/codeclimate/maintainability)

[![Test Coverage](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/test_coverage)](https://codeclimate.com/github/codeclimate/codeclimate/test_coverage)

# Payload Pro API

PayloadPro is a web application which gives you endpoints to send HTTP requests to and view the contents of the request.

It's primary purpose is for debugging connected application features, such as webhooks.

## URL Structure

https://api.payload.pro

https://api.payload.pro/bins

https://api.payload.pro/bins/{id}

https://api.payload.pro/bins/{id}/view

## Running locally

A docker compose file is available and you can bring up a stack with:

```
docker-compose up -d
```

This will create an API, the MongoDB 4.1.2 datastore and expose the API to you on `http://localhost:8081`

## Supports

 - Incoming JSON

## Todo

 - [ ] Fake response codes to test failures
 - [ ] Set a max input body size
