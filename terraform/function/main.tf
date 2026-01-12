resource "google_cloudfunctions2_function" "agents" {
  provider = google-beta
  name     = var.name
  location = var.region

  build_config {
    runtime     = "go125"
    entry_point = var.entrypoint
    source {
      storage_source {
        bucket = var.bucket
        object = var.object
      }
    }
  }

  service_config {
    max_instance_count    = 10
    available_memory      = "256Mi"
    timeout_seconds       = 60
    service_account_email = google_service_account.sa.email


    environment_variables = var.environment_variables

    direct_vpc_network_interface {
      network    = var.network
      subnetwork = var.subnetwork
    }
    direct_vpc_egress = "VPC_EGRESS_ALL_TRAFFIC"
  }
}

resource "google_service_account" "sa" {
  account_id = "${lower(var.name)}-sa"
}


resource "google_cloud_run_v2_service_iam_member" "public_access" {
  count = var.public_access ? 1 : 0

  name     = lower(google_cloudfunctions2_function.agents.name)
  location = google_cloudfunctions2_function.agents.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_project_iam_member" "accesses" {
  for_each = var.roles
  project  = google_cloudfunctions2_function.agents.project
  role     = each.key
  member   = google_service_account.sa.member
}
