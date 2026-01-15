module "function_job_offer" {
  source = "./function"


  environment_variables = {
    MCP_ENDPOINT = "${google_cloud_run_v2_service.mcp.uri}/sse"
  }

  bucket = google_storage_bucket.bucket.name
  object = google_storage_bucket_object.api.name

  network    = google_compute_network.vpc_network.name
  subnetwork = google_compute_subnetwork.public_subnet.name

  name       = "jobOfferAnalyzer"
  entrypoint = "JobOfferAnalyzer"

  public_access = false


  roles = ["roles/aiplatform.user"]
}

module "function_reachout" {
  source = "./function"

  bucket = google_storage_bucket.bucket.name
  object = google_storage_bucket_object.api.name

  network    = google_compute_network.vpc_network.name
  subnetwork = google_compute_subnetwork.public_subnet.name

  name       = "reachOutAgent"
  entrypoint = "ReachOutAgent"

  public_access = false


  roles = ["roles/aiplatform.user"]
}

module "function_cv_analyzer" {
  source = "./function"


  environment_variables = {
    MCP_ENDPOINT = "${google_cloud_run_v2_service.mcp.uri}/sse"
  }

  bucket = google_storage_bucket.bucket.name
  object = google_storage_bucket_object.api.name

  network    = google_compute_network.vpc_network.name
  subnetwork = google_compute_subnetwork.public_subnet.name

  name       = "cvAnalyzerAgent"
  entrypoint = "CvAnalyzerAgent"

  public_access = true

  roles = ["roles/aiplatform.user"]
}

