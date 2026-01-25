resource "google_cloud_run_v2_service" "mcp" {
  name     = "recrudb-mcp"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_INTERNAL_ONLY"

  template {
    timeout = "300s"
    service_account = google_service_account.cr_mcp_identity.email
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
      ports {
        container_port = 3000
      }

      resources {
        limits = {

          cpu    = "1"
          memory = "512Mi"
        }
      }

    }

    vpc_access {
      network_interfaces {
        network    = google_compute_network.vpc_network.name
        subnetwork = google_compute_subnetwork.private_subnet.name
      }
      egress = "ALL_TRAFFIC"
    }
  }

  deletion_protection = false

  lifecycle {
    ignore_changes = [
      template[0].containers[0].image,
    ]
  }

  depends_on = [google_project_service.gcp_services]
}


resource "google_service_account" "cr_mcp_identity" {
  account_id   = "recrudb-mcp-sa"
  display_name = "Cloud Run Firestore API Identity"
}

# Granting Firestore & Vertex AI access
resource "google_project_iam_member" "mcpm_roles" {
  for_each = toset(["roles/datastore.user"])
  project = google_service_account.cr_mcp_identity.project
  role    = each.key
  member  = google_service_account.cr_mcp_identity.member
}

resource "google_cloud_run_v2_service_iam_member" "cr_access" {
  for_each = toset([google_service_account.cr_api_identity.member])
  name     = google_cloud_run_v2_service.mcp.name
  location = google_cloud_run_v2_service.mcp.location
  role     = "roles/run.invoker"
  member   = each.key
}
