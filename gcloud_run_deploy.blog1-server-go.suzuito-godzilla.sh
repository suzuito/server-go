gcloud builds submit --config=./cloudbuild.server.suzuito-godzilla.yaml
gcloud run deploy blog1-server-go \
--image=gcr.io/suzuito-godzilla/server-go:latest \
--platform=managed \
--region=asia-northeast1 \
--allow-unauthenticated \
--memory=512Mi \
--cpu=1000m \
--max-instances=1 \
--set-env-vars="ENV=godzilla" \
--set-env-vars="PRERENDER_URL=https://prerendering-iflkmd4llq-an.a.run.app" \
--set-env-vars="FRONT_URL=https://suzuito-godzilla-blog1-server.storage.googleapis.com"