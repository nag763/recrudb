variable "region" {
  description = "Region to enable the feature"
  type        = string
  default     = "europe-west9"
}

variable "enabled_apis" {
  description = "List of GCP APIs to enable for the project"
  type        = list(string)
  default = [
    "aiplatform.googleapis.com",
    "analyticshub.googleapis.com",
    "artifactregistry.googleapis.com",
    "compute.googleapis.com",
    "cloudapis.googleapis.com",
    "cloudbuild.googleapis.com",
    "cloudfunctions.googleapis.com",
    "cloudtrace.googleapis.com",
    "containerregistry.googleapis.com",
    "datastore.googleapis.com",
    "firebaserules.googleapis.com",
    "firestore.googleapis.com",
    "iam.googleapis.com",
    "iamcredentials.googleapis.com",
    "logging.googleapis.com",
    "monitoring.googleapis.com",
    "run.googleapis.com",
    "servicemanagement.googleapis.com",
    "serviceusage.googleapis.com",
    "source.googleapis.com",
    "identitytoolkit.googleapis.com",
    "firebase.googleapis.com"
  ]
}