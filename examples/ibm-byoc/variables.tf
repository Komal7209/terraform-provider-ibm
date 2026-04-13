variable "bearer_token" {
  description = "Bearer token for BYOC API authentication"
  type        = string
  sensitive   = true
}

variable "subscription_id" {
  type        = string
  description = "BYOC subscription ID"
}

variable "dataplane_id" {
  type        = string
  description = "Dataplane ID (must be a valid UUID)"
}

variable "dataplane_name" {
  type        = string
  description = "Name for the BYOC dataplane"
  default     = "my-byoc-dataplane"
}

variable "region" {
  type        = string
  description = "IBM Cloud region where your BYOC dataplane will be created"
  default     = "us-east"
}

variable "cloud_provider" {
  type        = string
  description = "Cloud provider for the dataplane (aws, azure, or gcp)"
  default     = "aws"
  validation {
    condition     = contains(["aws", "azure", "gcp"], var.cloud_provider)
    error_message = "Cloud provider must be one of: aws, azure, gcp"
  }
}

# AWS-specific Dataplane Variables
variable "hyperscaler_account_id" {
  type        = string
  description = "AWS account ID (required for AWS deployments)"
  default     = null
}

variable "vpc_cidr" {
  type        = string
  description = "VPC CIDR block (required for AWS deployments)"
  default     = null
}

variable "vpc_network_type" {
  type        = string
  description = "VPC network type: public or private (required for AWS deployments)"
  default     = null
  validation {
    condition     = var.vpc_network_type == null || contains(["public", "private"], var.vpc_network_type)
    error_message = "VPC network type must be either 'public' or 'private'"
  }
}

variable "private_subnet_cidrs" {
  type        = list(string)
  description = "Private subnet CIDR blocks (optional for AWS deployments)"
  default     = null
}

variable "public_subnet_cidrs" {
  type        = list(string)
  description = "Public subnet CIDR blocks (optional for AWS deployments)"
  default     = null
}

# Azure-specific Dataplane Variables
variable "hyperscaler_subscription_id" {
  type        = string
  description = "Azure subscription ID (required for Azure deployments)"
  default     = null
}

variable "hyperscaler_tenant_id" {
  type        = string
  description = "Azure tenant ID (required for Azure deployments)"
  default     = null
}

# Engine Variables
variable "engine_name" {
  type        = string
  description = "Name for the BYOC engine"
  default     = "my-db2-engine"
}

variable "engine_type" {
  type        = string
  description = "Type of database engine (db2)"
  default     = "db2"
  validation {
    condition     = contains(["db2"], var.engine_type)
    error_message = "Engine type must be 'db2'"
  }
}

# Optional Engine Configuration Variables
variable "availability_zone" {
  type        = string
  description = "Availability zone for the engine"
  default     = null
}

variable "storage_units" {
  type        = number
  description = "Number of storage units for the engine"
  default     = null
}

variable "compute_units" {
  type        = number
  description = "Number of compute units for the engine"
  default     = null
}

variable "instance_type" {
  type        = string
  description = "Instance type for the engine"
  default     = null
}

variable "endpoint_type" {
  type        = string
  description = "Endpoint type: private, public, or public-private"
  default     = null
  validation {
    condition     = var.endpoint_type == null || contains(["private", "public", "public-private"], var.endpoint_type)
    error_message = "Endpoint type must be one of: private, public, public-private"
  }
}

variable "private_link_service_enabled" {
  type        = bool
  description = "Enable private link service for the engine"
  default     = null
}

variable "public_enabled" {
  type        = bool
  description = "Enable public access for the engine"
  default     = null
}

variable "oracle_compatibility" {
  type        = bool
  description = "Enable Oracle compatibility mode"
  default     = null
}

variable "replicas" {
  type        = number
  description = "Number of replicas for the engine"
  default     = null
}

variable "plan" {
  type        = string
  description = "Service plan for the engine"
  default     = null
}

variable "profile_name" {
  type        = string
  description = "Profile name for the engine"
  default     = null
}