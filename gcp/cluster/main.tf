resource "google_container_cluster" "primary" {
  name               = "${var.cluster_name}"
  project            = "${var.project_id}"
  zone               = "${var.master_zone}"
  initial_node_count = 1
  node_version = "1.7.8-gke.0"

  master_auth {
    username = "${var.master_username}"
    password = "${var.master_password}"
  }
  
  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]

    machine_type = "n1-standard-1"
  }
}
