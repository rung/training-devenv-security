terraform {
  backend "gcs" {
    # You need to modify bucket
    bucket = "tmp-training-devenv-security-rung"
    prefix = "terraform/state"
  }
}
