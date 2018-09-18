# Payload Pro API

[![Maintainability](https://api.codeclimate.com/v1/badges/a2b86c9814643d6cc55a/maintainability)](https://codeclimate.com/github/PayloadPro/pro.payload.api/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a2b86c9814643d6cc55a/test_coverage)](https://codeclimate.com/github/PayloadPro/pro.payload.api/test_coverage)

The PayloadPro API container is a Docker container that contains a Go binary built from this repo for running the PayloadPro API. It exposes API endpoints to interact with the bins, requests and stats.

You need to back the API with a Cockroach DB instance, which if you're looking to just run this locally (or publicly with your own resources), you should use the [ops](https://github.com/PayloadPro/ops) repository where this container (and other dependencies) can be ran with a simple `docker-compose up -d` command.

## Environment

When running the container you can use the following environment variables:

| ENV             | Description                           | Default
| --------------- | ------------------------------------- | -------
| `APP_NAME`      | The name the app presents itself with | `Payload Pro`
| `APP_SITE_LINK` | A link to the website                 | `http://localhost:3000`
| `APP_API_LINK`  | A link to the API                     | `http://localhost:8081`
| `APP_DOCS_LINK` | A link to the docs                    | `http://localhost:3000/rtfm`
| `DB_DSN`        | CockroachDB compatible DSN            | `postgresql://pp@cockroach-proxy:26257/payloadpro?sslmode=disable`

PayloadPro is a web application which gives you endpoints to send HTTP requests to and view the contents of the request.

It's primary purpose is for debugging connected application features, such as webhooks.
