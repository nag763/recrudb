resource "google_cloud_run_v2_service" "mcp" {
  name     = "recrudb-mcp"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_INTERNAL_ONLY"
  provider = google-beta

  template {

    containers {
      image          = "scratch"
      base_image_uri = "us-central1-docker.pkg.dev/serverless-runtimes/google-22/runtimes/go125"
      command        = ["./mcp"]
      ports {
        container_port = 3000
      }

      resources {
        limits = {

          cpu    = "1"
          memory = "512Mi"
        }
      }

      source_code {
        cloud_storage_source {
          bucket     = google_storage_bucket.bucket.name
          object     = google_storage_bucket_object.mcp.name
          generation = google_storage_bucket_object.mcp.generation
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

resource "google_cloud_run_v2_service_iam_member" "cr_access" {
  for_each = toset([ module.function_job_offer.service_account.member, module.function_cv_analyzer.service_account.member ])
  name     = google_cloud_run_v2_service.mcp.name
  location = google_cloud_run_v2_service.mcp.location
  role     = "roles/run.invoker"
  member   = each.key
}
