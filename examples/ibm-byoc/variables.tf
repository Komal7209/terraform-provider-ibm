variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "byoc_bearer_token" {   
    description = "Bearer token for BYOC API"
    type        = string
}

// Resource arguments for byoc_engine
variable "byoc_engine_subscription_id" {
  description = "Subscription ID."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}
variable "byoc_engine_dataplane_id" {
  description = "Data Plane ID."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}
variable "byoc_engine_availability_zone" {
  description = ""
  type        = string
  default     = "us-south-1"
}
variable "byoc_engine_storage_units" {
  description = ""
  type        = number
  default     = 50
}
variable "byoc_engine_compute_units" {
  description = ""
  type        = number
  default     = 2
}
variable "byoc_engine_engine_name" {
  description = ""
  type        = string
  default     = "db2-engine"
}
variable "byoc_engine_engine_type" {
  description = ""
  type        = string
  default     = "db2"
}
variable "byoc_engine_endpoint_type" {
  description = ""
  type        = string
  default     = "private"
}
variable "byoc_engine_instance_type" {
  description = ""
  type        = string
  default     = "Standard_D4s_v5"
}
variable "byoc_engine_replicas" {
  description = ""
  type        = number
  default     = 1
}
variable "byoc_engine_public_enabled" {
  description = ""
  type        = bool
  default     = false
}
variable "byoc_engine_private_link_service_enabled" {
  description = ""
  type        = bool
  default     = true
}
variable "byoc_engine_oracle_compatibility" {
  description = ""
  type        = bool
  default     = false
}
variable "byoc_engine_plan" {
  description = ""
  type        = string
  default     = "db2wh-small"
}
variable "byoc_engine_profile_name" {
  description = ""
  type        = string
  default     = "netezza-profile"
}
variable "byoc_engine_service_principals" {
  description = ""
  type        = list(string)
  default     = ["service-principal-1"]
}
variable "byoc_engine_subscription_ids" {
  description = ""
  type        = list(string)
  default     = ["550e8400-e29b-41d4-a716-446655440000"]
}
