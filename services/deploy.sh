#!/bin/sh

gcloud run deploy recrudb-api   --source .   --region europe-west9   --set-build-env-vars GOOGLE_BUILDABLE=./cmd/api & gcloud run deploy recrudb-mcp   --source .   --region europe-west9   --set-build-env-vars GOOGLE_BUILDABLE=./cmd/mcp
