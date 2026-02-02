resource "google_firebase_hosting_site" "recrudb_front" {
  provider = google-beta
  site_id  = "recrudb"
}

resource "google_firebase_hosting_custom_domain" "recrudb_front" {
  provider      = google-beta
  site_id       = google_firebase_hosting_site.recrudb_front.site_id
  custom_domain = "www.recrudb.xyz"
}