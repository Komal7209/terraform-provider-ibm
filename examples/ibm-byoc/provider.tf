provider "ibm" {
  # BYOC API uses a dedicated bearer token and must not be routed through iam_token
  byoc_bearer_token = var.bearer_token
  region            = var.region
}