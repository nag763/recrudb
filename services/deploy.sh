#!/usr/bin/env bash
set -euo pipefail

REGION="europe-west9"
SOURCE="."
BUILD_VAR="GOOGLE_BUILDABLE"

deploy() {
  local SERVICE_NAME=$1
  local BUILD_PATH=$2

  echo "🚀 Deploying ${SERVICE_NAME} (${BUILD_PATH})..."

  gcloud run deploy "${SERVICE_NAME}" \
    --source "${SOURCE}" \
    --region "${REGION}" \
    --set-build-env-vars "${BUILD_VAR}=${BUILD_PATH}"
}

TARGET="${1:-all}"

case "$TARGET" in
  api)
    deploy recrudb-api ./cmd/api
    ;;
  mcp)
    deploy recrudb-mcp ./cmd/mcp
    ;;
  all)
    deploy recrudb-api ./cmd/api &
    deploy recrudb-mcp ./cmd/mcp &
    wait
    ;;
  *)
    echo "❌ Unknown target: $TARGET"
    echo "Usage: $0 [api|mcp]"
    exit 1
    ;;
esac

echo "✅ Deployment finished"
