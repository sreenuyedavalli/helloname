resource "google_sql_database_instance" "master" {
  name             = "${var.db_instance_name}"
  project          = "${var.project_id}"
  region           = "${var.region}"
  database_version = "POSTGRES_9_6"
  settings {
    tier = "db-f1-micro"
    disk_size = "10"
    backup_configuration {
      enabled = "false"
    }
  }
}

resource "google_sql_database" "nyt-nametrack" {
  name      = "namecount"
  instance  = "${google_sql_database_instance.master.name}"
}

resource "google_sql_user" "primary-db-user" {
  name      = "dbadmin"
  instance  = "${google_sql_database_instance.master.name}"
  host      = ""
  password  = "${var.primary_db_user_password}"
}
