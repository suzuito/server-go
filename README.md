SPA server

- In case of bot request, send the request to a prerender.
- In case of not bot request, send the request to a server delivering SPA.

## Developments

```bash
gcloud auth login
gcloud auth configure-docker
```

```bash
docker-compose up
```

```bash
make test
```

### blog1-server:8080

```hosts
127.0.0.1 blog1-server
```

```bash
open http://blog1-server:8080
```

Prerenderingの検証。

```bash
curl -A "googlebot" http://blog1-server:8080 > a.html
open a.html
```
