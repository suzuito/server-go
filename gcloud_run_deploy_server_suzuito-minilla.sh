
gcloud run deploy blog1-server-go \
--image=gcr.io/${env.GCP_PROJECT_ID}/blog1-server-go:latest \
--platform=managed \
--region=asia-northeast1 \
--allow-unauthenticated \
--memory=512Mi \
--cpu=1000m \
--max-instances=1 \
--set-env-vars="ENV=dev" \
--set-env-vars="PRERENDER_URL=https://prerendering-fgfw3lzk2a-an.a.run.app" \
--set-env-vars="FRONT_URL=https://suzuito-minilla-blog1-server.storage.googleapis.com"