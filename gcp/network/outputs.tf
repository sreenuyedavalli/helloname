
output "NGINX IP:" {
  value = "${google_compute_address.nginx-ip.address}"
}
