resource "google_service_account" "gh_wif_identity" {
  account_id   = "recrudb-wif-sa"
  display_name = "Github WIF Service Account"
}

resource "google_iam_workload_identity_pool" "github_pool" {
  workload_identity_pool_id = "github-actions-pool"
  display_name              = "GitHub Actions Pool"
}

resource "google_iam_workload_identity_pool_provider" "github_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-provider"
  display_name                       = "GitHub OIDC Provider"
  description                        = "OIDC Provider for GitHub Actions"
  attribute_condition                = "attribute.repository == 'nag763/recrudb'"
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.actor"      = "assertion.actor"
    "attribute.aud"        = "assertion.aud"
    "attribute.repository" = "assertion.repository"
  }
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
  disabled = false
}

resource "google_project_iam_member" "wif_roles" {
  for_each = toset(
    ["roles/run.admin",
      "roles/firebasehosting.admin",
      "roles/cloudbuild.builds.editor",
      "roles/artifactregistry.writer",
      "roles/iam.serviceAccountUser",
      "roles/storage.admin",
      "roles/serviceusage.serviceUsageConsumer"
  ])
  project = google_service_account.gh_wif_identity.project
  role    = each.key
  member  = google_service_account.gh_wif_identity.member
}


# Allow the GitHub Repo to impersonate your Firebase Service Account
resource "google_service_account_iam_member" "wif_binding" {
  service_account_id = google_service_account.gh_wif_identity.id
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/projects/${data.google_project.project.number}/locations/global/workloadIdentityPools/${google_iam_workload_identity_pool.github_pool.workload_identity_pool_id}/attribute.repository/nag763/recrudb"
}
