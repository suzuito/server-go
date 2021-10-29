

## Developments

```bash
gcloud auth login
gcloud auth configure-docker
```

```bash
docker-compose up
```

```bash
docker-compose exec blog1-server /bin/bash -c 'make test'
```

### blog1-server:8080

以下の行をhostsファイルへ追加する

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
