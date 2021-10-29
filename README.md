

## Developments

```bash
gcloud auth login
gcloud auth configure-docker
```

```bash
docker-compose up
```

以下の行をhostsファイルへ追加する

```hosts
127.0.0.1 blog1-server
```

### blog1-server:8080

```bash
open http://blog1-server:8080
```

Prerenderingの検証。

```bash
curl -A "googlebot" http://blog1-server:8080 > a.html
open a.html
```
