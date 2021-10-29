

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
127.0.0.1 server
```

検証用URLは http://server:8080 である。

```bash
open http://server:8080
```

Prerenderingの検証。

```bash
curl -A "googlebot" http://server:8080 > a.html
open a.html
```
