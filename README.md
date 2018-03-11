# Envoy Demo

## Description

App to demo a few basic usage patterns [Envoy][0].

You can cross-reference the Envoy configs in this project with the [v2 config docs][4].

## Requirements

You should be able to run [docker-compose version 3][1].

## Sample Usage Patterns

This list is a work-in-progress.

### Simple Envoy sidecar setup

Run a service and put an Envoy proxy in front of it. In addition to routing requests, the Envoy instance will periodically ping the service for health status.

The Envoy configuration can be found [here][2].

```bash
docker-compose -f docker-compose.simple-proxy.yaml up --build
# wait till service is ready

# From terminal:
curl localhost:3000/health
curl localhost:3000/info

# From browser, visit the Envoy admin page:
GET localhost:3001

# From browser, visit the Envoy stats page:
GET localhost:3001/stats
```

### Sample Request Tracing using Zipkin

Run the same service as before, do a few jumps to other endpoints. Visit the Zipkin dashboard to view the request traces.

The Envoy configuration can be found [here][3].

```bash
docker-compose -f docker-compose.request-tracing.yaml up --build
# wait till service is ready

# Make the same requests as before and/or:
curl localhost:3000/jump

# Visit the Zipkin dashboard from your browser:
GET localhost:9411/
```

[0]: https://www.envoyproxy.io/
[1]: https://docs.docker.com/compose/compose-file/
[2]: ./envoy/simple-proxy.yaml
[3]: ./envoy/request-tracing.yaml
[4]: https://www.envoyproxy.io/docs/envoy/latest/api-v2/api
