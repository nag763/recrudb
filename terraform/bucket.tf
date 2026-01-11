resource "google_storage_bucket" "bucket" {
  name                        = "recrudb-deploy-source"
  location                    = var.region
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "api" {
  name   = "source.zip#${data.archive_file.api_function_zip.output_md5}"
  bucket = google_storage_bucket.bucket.name
  source = data.archive_file.api_function_zip.output_path
}

resource "google_storage_bucket_object" "mcp" {
  name   = "source.zip#${data.archive_file.mcp_function_zip.output_md5}"
  bucket = google_storage_bucket.bucket.name
  source = data.archive_file.mcp_function_zip.output_path
}