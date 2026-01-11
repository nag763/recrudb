provider "google" {
  project = "recrudb"
  region  = var.region
}

provider "google-beta" {
  project = "recrudb"
}


resource "google_project_service" "gcp_services" {
  for_each = toset(var.enabled_apis)
  service  = each.key

  disable_dependent_services = true
}

resource "google_artifact_registry_repository" "recrudb" {
  repository_id = "cloud-run-source-deploy"
  format        = "DOCKER"

  depends_on = [google_project_service.gcp_services]
}