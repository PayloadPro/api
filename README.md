# Payload Pro API

[![Maintainability](https://api.codeclimate.com/v1/badges/a2b86c9814643d6cc55a/maintainability)](https://codeclimate.com/github/PayloadPro/pro.payload.api/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a2b86c9814643d6cc55a/test_coverage)](https://codeclimate.com/github/PayloadPro/pro.payload.api/test_coverage)

PayloadPro is a web application which gives you endpoints to send HTTP requests to and view the contents of the request.

It's primary purpose is for debugging connected application features, such as webhooks.

## URL Structure

[ GET ] https://api.payload.pro

[ GET|POST ] https://api.payload.pro/bins

[ GET ] https://api.payload.pro/bins/{id}

[ ANY ] https://api.payload.pro/bins/{id}/request

[ GET ] https://api.payload.pro/bins/{id}/requests

[ GET ] https://api.payload.pro/bins/{id}/requests/{req_id}

## Running locally

A docker compose file is available and you can bring up a stack with:

```
docker-compose up -d
```

This will create an API, the MongoDB 4.1.2 datastore and expose the API to you on `http://localhost:8081`

## Supports

 - Incoming JSON

## Todo

 - [ ] Automatically setup database and collections
 - [ ] Proxy methods to forward incoming webhooks to enable MITM debugging
 - [ ] Fake responses to test failure scenarios
 - [ ] Set a max input body size for public API
 - [ ] Create a public docker hub image
