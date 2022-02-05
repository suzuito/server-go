SPA server

- In case of bot request, send the request to a prerenderer.
- In case of not bot request, send the request to a server delivering SPA.

## Developments

```bash
docker-compose up
```

```bash
make test
```

```bash
# The health check
curl -A "kube-probe" http://localhost:8080 -v
# The request is sent to a prerenderer (defined by PRERENDER_URL) and prernderer responses http://example.com page
curl -A "googlebot" -H "Host: example.com" http://localhost:8080 -v
# The request is sent to a front_end (defined by FRONT_URL)
curl http://localhost:8080 -v
```

## Operation

```bash
kubectl create namespace blog
kubectl apply -k kustomize/blog/minilla/bases
kubectl rollout restart deployment front-deployment
kubectl logs -l app=front -c front
kubectl logs -l app=front -c prerendering
```