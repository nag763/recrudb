resource "github_actions_secret" "wif_provider" {
  repository      = "recrudb"
  secret_name     = "WIF_PROVIDER"
  plaintext_value = google_iam_workload_identity_pool_provider.github_provider.name
}

# The Service Account email
resource "github_actions_secret" "wif_service_account" {
  repository      = "recrudb"
  secret_name     = "WIF_SERVICE_ACCOUNT"
  plaintext_value = google_service_account.gh_wif_identity.email
}

resource "github_actions_secret" "firebase_api_key" {
  repository = "recrudb"
  secret_name = "FIREBASE_API_KEY"
  plaintext_value = var.firebase_api_key
}

resource "github_actions_variable" "region" {
  repository = "recrudb"
  variable_name = "GCP_REGION"
  value = var.region
}