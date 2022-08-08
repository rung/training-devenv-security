# Configuration for helloworld application
resource "google_service_account" "run-helloworld" {
  account_id = "run-helloworld"
}

resource "google_artifact_registry_repository" "training-helloworld" {
  location      = "us"
  repository_id = "training-helloworld"
  format        = "DOCKER"
}
