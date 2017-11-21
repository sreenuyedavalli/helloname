module "cluster" {
  source = "cluster/"

  project_id      = "${var.project_id}"
  cluster_name    = "${var.cluster_name}"
  master_password = "${var.master_password}"
}

module "network" {
  source = "network/"

  project_id = "${var.project_id}"
}

module "sql" {
  source = "sql/"

  project_id               = "${var.project_id}"
  region                   = "${var.region}"
  db_instance_name         = "${var.db_instance_name}"
  primary_db_user_password = "${var.primary_db_user_password}"
}
