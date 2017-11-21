# Reserved IPs

resource "google_compute_address" "nginx-ip" {
  project = "${var.project_id}"
  name    = "nginx-static-ip"
}
