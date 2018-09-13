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

This will create:

 - API on `http://localhost:8081`
 - CockroachDB UI on `http://localhost:8080`
 - CockroachDB cluster on `http://localhost:26257`


## Todo

 - [ ] Proxy methods to forward incoming webhooks to enable MITM debugging
 - [ ] Fake responses to test failure scenarios
 - [ ] Set a max input body size for public API
 - [ ] Create a public docker hub image


## Creating HA Proxy Configs

Form you local machine, you can connect to the cluster via `localhost` so you can run:

```
cd ./deployments && \
cockroach gen haproxy --insecure --host localhost
```

This will generate the HA Proxy config based on discovery with the cluster. You can change the host and connection methods, as well as the output file depending ont he environment you're configuring.


## Load Testing

A basic load test is available in `./tests/load/`. You can run it with:

```
./tests/load/load.sh
```
