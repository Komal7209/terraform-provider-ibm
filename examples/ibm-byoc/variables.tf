variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}
// Resource arguments for byoc_engine
variable "byoc_engine_subscription_id" {
  description = "Subscription ID"
  type        = string
  default     = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
}
variable "byoc_engine_dataplane_id" {
  description = "Dataplane ID"
  type        = string
  default     = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
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
  description = "Type of the engine (db2, db2wh, netezza)"
  type        = string
  default     = "db2"
}
variable "byoc_engine_admin_username" {
  description = "Admin username for the engine"
  type        = string
  default     = "admin"
}
variable "byoc_engine_admin_password" {
  description = "Admin password for the engine (SHA-2 hashed)"
  type        = string
  default     = "{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA=="
  sensitive   = true
}
variable "byoc_engine_admin_email" {
  description = "Admin email for the engine"
  type        = string
  default     = "admin@example.com"
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
