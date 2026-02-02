resource "google_cloud_run_v2_service" "api" {
  name     = "recrudb-api"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    timeout = "90s"
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
      ports {
        container_port = 8000
      }


      resources {
        limits = {

          cpu    = "1"
          memory = "512Mi"
        }
      }

    }
    service_account = google_service_account.cr_api_identity.email
  }

  deletion_protection = false

  lifecycle {
    ignore_changes = [
      template[0].containers[0].image,
    ]
  }

  depends_on = [google_project_service.gcp_services]
}

resource "google_service_account" "cr_api_identity" {
  account_id   = "recrudb-api-sa"
  display_name = "Cloud Run Firestore API Identity"
}

# Granting Firestore & Vertex AI access
resource "google_project_iam_member" "roles" {
  for_each = toset(["roles/aiplatform.user", "roles/datastore.user"])
  project = google_service_account.cr_api_identity.project
  role    = each.key
  member  = google_service_account.cr_api_identity.member
}

resource "google_cloud_run_v2_service_iam_member" "api_cr_access" {
  for_each = toset(["allUsers"])
  name     = google_cloud_run_v2_service.api.name
  location = google_cloud_run_v2_service.api.location
  role     = "roles/run.invoker"
  member   = each.key
}
