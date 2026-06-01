provider "ibm" {
  # BYOC API uses a dedicated bearer token and must not be routed through iam_token
  # Set the bearer token using environment variable: BYOC_BEARER_TOKEN or IBMCLOUD_BYOC_TOKEN
  # Example: export BYOC_BEARER_TOKEN="your-bearer-token-here"
  ibmcloud_api_key = var.ibmcloud_api_key
}