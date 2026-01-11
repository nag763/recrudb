resource "google_compute_network" "vpc_network" {
  name                    = "recrudb-vpc"
  auto_create_subnetworks = false # Crucial: This enables Custom Mode
  routing_mode            = "REGIONAL"


  depends_on = [google_project_service.gcp_services]
}

resource "google_compute_subnetwork" "public_subnet" {
  name          = "public"
  ip_cidr_range = "10.0.1.0/24"
  region        = var.region
  network       = google_compute_network.vpc_network.id

  private_ip_google_access = true
}

resource "google_compute_subnetwork" "private_subnet" {
  name          = "private"
  ip_cidr_range = "10.0.2.0/24"
  region        = var.region
  network       = google_compute_network.vpc_network.id

  private_ip_google_access = true
}

resource "google_compute_firewall" "mcp_fw" {
  name    = "allow-https-inbound"
  network = google_compute_network.vpc_network.name

  allow {
    protocol = "tcp"
    ports    = ["443"]
  }

  source_ranges = [google_compute_subnetwork.public_subnet.ip_cidr_range, google_compute_subnetwork.private_subnet.ip_cidr_range]
}