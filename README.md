[![codecov](https://codecov.io/gh/suzuito/server-go/branch/main/graph/badge.svg?token=9ABQLNMSDV)](https://codecov.io/gh/suzuito/server-go)

SPA server

- In case of bot request, send the request to a prerenderer.
- In case of not bot request, send the request to a server delivering SPA.

## Developments

```bash
docker-compose up
```

Generate mock

```bash
sh mockgen.sh
```

Run test

```bash
make test
```

```bash
# The health check
curl -A "kube-probe" http://localhost:8080 -v
# The request is sent to a prerenderer (defined by PRERENDER_URL) and prernderer responses http://example.com page
curl -A "googlebot" -H "Host: suzuito.github.io" http://localhost:8080/test-pages/page001.html -v
# The request is sent to a front_end (defined by FRONT_URL https://suzuito.github.io/test-pages)
curl http://localhost:8080/page001.html -v
curl http://localhost:8080/main.js -v
# Sitemap
curl http://localhost:8080/sitemap.xml -v
```

## Deployment

- CloudBuild
