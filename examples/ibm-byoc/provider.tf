provider "ibm" {
  # BYOC API uses bearer token authentication
  iam_token = var.bearer_token
  region    = var.region
}