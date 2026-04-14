# ============================================================================
# IBM BYOC (Bring Your Own Cloud) Terraform Example
# ============================================================================
# This example demonstrates CRUD operations for BYOC dataplanes and engines
# Supports: AWS, Azure, GCP
# Engine Types: Db2, Db2 Warehouse, Netezza
# ============================================================================

# ============================================================================
# DATAPLANE CONFIGURATION
# ============================================================================
# Uncomment the section for your cloud provider (AWS or Azure)
# Only one dataplane configuration should be active at a time

# ----------------------------------------------------------------------------
# Option 1: AWS Dataplane
# ----------------------------------------------------------------------------
# resource "ibm_byoc_dataplane" "byoc_dataplane" {
#   subscription_id = var.subscription_id
#   dataplane_id    = var.dataplane_id
#   name            = var.dataplane_name
#   region          = var.aws_region  # e.g., "us-east-1"
#   cloud_provider  = "AWS"
#
#   # AWS-specific required fields
#   aws_account_id = var.aws_account_id
#   aws_role_arn   = var.aws_role_arn
#
#   # AWS-specific optional network configuration
#   vpc_cidr             = var.vpc_cidr              # e.g., "10.0.0.0/16"
#   vpc_network_type     = var.vpc_network_type      # e.g., "private" or "public"
#   private_subnet_cidrs = var.private_subnet_cidrs  # e.g., ["10.0.1.0/24", "10.0.2.0/24"]
#   public_subnet_cidrs  = var.public_subnet_cidrs   # e.g., ["10.0.101.0/24", "10.0.102.0/24"]
#
#   timeouts {
#     create = "30m"
#     update = "30m"
#     delete = "30m"
#   }
# }

# ----------------------------------------------------------------------------
# Option 2: Azure Dataplane with Db2 Engine (ACTIVE FOR TESTING)
# ----------------------------------------------------------------------------
# resource "ibm_byoc_dataplane" "byoc_dataplane" {
#   subscription_id = var.subscription_id
#   dataplane_id    = var.dataplane_id
#   name            = var.dataplane_name
#   region          = var.azure_region  # e.g., "eastus"
#   cloud_provider  = "Azure"

#   # Azure-specific required fields
#   hyperscaler_subscription_id = var.azure_subscription_id
#   hyperscaler_tenant_id       = var.azure_tenant_id

#   timeouts {
#     create = "30m"
#     update = "30m"
#     delete = "30m"
#   }
# }

# ============================================================================
# ENGINE CONFIGURATION
# ============================================================================
# Uncomment the engine type you want to deploy
# Multiple engines can be deployed on the same dataplane

# ----------------------------------------------------------------------------
# Option 1: Db2 Engine (ACTIVE FOR TESTING)
# ----------------------------------------------------------------------------
# resource "ibm_byoc_engine" "db2_engine" {
#   subscription_id = var.subscription_id
#   dataplane_id    = ibm_byoc_dataplane.byoc_dataplane.dataplane_id
#   engine_name     = var.engine_name
#   engine_type     = "db2"

#   # Resource configuration
#   availability_zone = var.availability_zone  # e.g., "eastus-1"
#   storage_units     = var.storage_units      # e.g., 100
#   compute_units     = var.compute_units      # e.g., 4
#   instance_type     = var.instance_type      # e.g., "Standard_D4s_v3"

#   # Network configuration
#   endpoint_type                = var.endpoint_type                 # e.g., "private"
#   private_link_service_enabled = var.private_link_service_enabled  # e.g., true
#   public_enabled               = var.public_enabled                # e.g., false

#   # Additional configuration
#   oracle_compatibility = var.oracle_compatibility  # e.g., false
#   replicas             = var.replicas              # e.g., 1
#   plan                 = var.plan                  # e.g., "standard"
#   profile_name         = var.profile_name          # e.g., "db2-standard"

#   timeouts {
#     create = "60m"
#     update = "60m"
#     delete = "60m"
#   }

#   depends_on = [ibm_byoc_dataplane.byoc_dataplane]
# }

# ----------------------------------------------------------------------------
# Option 2: Db2 Warehouse Engine
# ----------------------------------------------------------------------------
# resource "ibm_byoc_engine" "db2wh_engine" {
#   subscription_id = var.subscription_id
#   dataplane_id    = ibm_byoc_dataplane.byoc_dataplane.dataplane_id
#   engine_id       = var.db2wh_engine_id
#   engine_name     = "${var.dataplane_name}-db2wh"
#   engine_type     = "db2wh"
#
#   # Resource configuration
#   availability_zone = var.availability_zone
#   storage_units     = 200
#   compute_units     = 8
#   instance_type     = var.instance_type
#
#   # Network configuration
#   endpoint_type  = "private"
#   public_enabled = false
#
#   timeouts {
#     create = "60m"
#     update = "60m"
#     delete = "60m"
#   }
#
#   depends_on = [ibm_byoc_dataplane.byoc_dataplane]
# }

# ----------------------------------------------------------------------------
# Option 3: Netezza Engine
# ----------------------------------------------------------------------------
# resource "ibm_byoc_engine" "netezza_engine" {
#   subscription_id = var.subscription_id
#   dataplane_id    = ibm_byoc_dataplane.byoc_dataplane.dataplane_id
#   engine_id       = var.netezza_engine_id
#   engine_name     = "${var.dataplane_name}-netezza"
#   engine_type     = "netezza"
#
#   # Resource configuration
#   availability_zone = var.availability_zone
#   storage_units     = 500
#   compute_units     = 16
#   instance_type     = var.instance_type
#
#   # Network configuration
#   endpoint_type  = "private"
#   public_enabled = false
#
#   timeouts {
#     create = "60m"
#     update = "60m"
#     delete = "60m"
#   }
#
#   depends_on = [ibm_byoc_dataplane.byoc_dataplane]
# }

# ============================================================================
# DATA SOURCES (for READ operations)
# ============================================================================
# Use data sources to query existing BYOC resources without managing them

# ----------------------------------------------------------------------------
# Example 1: Query a specific dataplane by ID
# ----------------------------------------------------------------------------
# Use this to get details about an existing dataplane
data "ibm_byoc_dataplane" "dataplane_info" {
  subscription_id = var.subscription_id
  dataplane_id    = var.dataplane_id #ibm_byoc_dataplane.byoc_dataplane.dataplane_id

  # depends_on = [ibm_byoc_dataplane.byoc_dataplane]
}

# ----------------------------------------------------------------------------
# Example 2: List all dataplanes in a subscription
# ----------------------------------------------------------------------------
# Use this to discover all dataplanes for your subscription
# data "ibm_byoc_dataplanes" "all_dataplanes" {
#   subscription_id = var.subscription_id

#   depends_on = [ibm_byoc_dataplane.byoc_dataplane]
# }

# ----------------------------------------------------------------------------
# Example 3: Query a specific Db2 engine by ID
# ----------------------------------------------------------------------------
# Use this to get details about an existing Db2 engine
# data "ibm_byoc_engine" "db2_engine_info" {
#   subscription_id = var.subscription_id
#   dataplane_id    = ibm_byoc_dataplane.byoc_dataplane.dataplane_id
#   engine_id       = ibm_byoc_engine.db2_engine.engine_id

#   depends_on = [ibm_byoc_engine.db2_engine]
# }

# ----------------------------------------------------------------------------
# Example 4: List all engines on a dataplane
# ----------------------------------------------------------------------------
# Use this to discover all engines deployed on a specific dataplane
# data "ibm_byoc_engines" "all_engines" {
#   subscription_id = var.subscription_id
#   dataplane_id    = ibm_byoc_dataplane.byoc_dataplane.dataplane_id

#   depends_on = [ibm_byoc_engine.db2_engine]
# }

# ----------------------------------------------------------------------------
# Example 5: Query existing resources (without creating them)
# ----------------------------------------------------------------------------
# Uncomment to query existing resources without managing them with Terraform
# This is useful for importing existing infrastructure or read-only access

# # Query an existing dataplane (not managed by this Terraform config)
# data "ibm_byoc_dataplane" "existing_dataplane" {
#   subscription_id = "your-subscription-id"
#   dataplane_id    = "existing-dataplane-id"
# }
#
# # Query an existing Db2 engine (not managed by this Terraform config)
# data "ibm_byoc_engine" "existing_db2" {
#   subscription_id = "your-subscription-id"
#   dataplane_id    = "existing-dataplane-id"
#   engine_id       = "existing-engine-id"
# }
#
# # List all dataplanes to discover what's deployed
# data "ibm_byoc_dataplanes" "discover_all" {
#   subscription_id = "your-subscription-id"
# }
#
# # List all engines on a specific dataplane
# data "ibm_byoc_engines" "discover_engines" {
#   subscription_id = "your-subscription-id"
#   dataplane_id    = "existing-dataplane-id"
# }

# ============================================================================
# OUTPUTS - Display information from data sources
# ============================================================================

# Output dataplane information
output "dataplane_details" {
  description = "Details of the BYOC dataplane"
  value = {
    id             = data.ibm_byoc_dataplane.dataplane_info.id
    name           = data.ibm_byoc_dataplane.dataplane_info.name
    cloud_provider = data.ibm_byoc_dataplane.dataplane_info.cloud_provider
    region         = data.ibm_byoc_dataplane.dataplane_info.region
    status         = data.ibm_byoc_dataplane.dataplane_info.status
    created_at     = data.ibm_byoc_dataplane.dataplane_info.created_at
  }
}

# Output all dataplanes in subscription
# output "all_dataplanes" {
#   description = "List of all dataplanes in the subscription"
#   value = [
#     for dp in data.ibm_byoc_dataplanes.all_dataplanes.dataplanes : {
#       id     = dp.id
#       name   = dp.name
#       status = dp.status
#       region = dp.region
#       cloud  = dp.cloud_provider
#     }
#   ]
# }

# Output Db2 engine information
# output "db2_engine_details" {
#   description = "Details of the Db2 engine"
#   value = {
#     id         = data.ibm_byoc_engine.db2_engine_info.id
#     name       = data.ibm_byoc_engine.db2_engine_info.name
#     type       = data.ibm_byoc_engine.db2_engine_info.type
#     status     = data.ibm_byoc_engine.db2_engine_info.status
#     version    = data.ibm_byoc_engine.db2_engine_info.version
#     endpoint   = data.ibm_byoc_engine.db2_engine_info.endpoint
#     created_at = data.ibm_byoc_engine.db2_engine_info.created_at
#   }
# }

# Output all engines on the dataplane
# output "all_engines" {
#   description = "List of all engines on the dataplane"
#   value = [
#     for engine in data.ibm_byoc_engines.all_engines.engines : {
#       id       = engine.id
#       name     = engine.name
#       type     = engine.type
#       status   = engine.status
#       endpoint = engine.endpoint
#     }
#   ]
# }

# ============================================================================
# CRUD OPERATIONS GUIDE
# ============================================================================
#
# CREATE:
#   terraform apply
#
# READ (using data sources):
#   terraform show
#   terraform output dataplane_details
#   terraform output db2_engine_details
#   terraform output all_dataplanes
#   terraform output all_engines
#
# READ (using state):
#   terraform state show ibm_byoc_dataplane.byoc_dataplane
#   terraform state show ibm_byoc_engine.db2_engine
#   terraform state show data.ibm_byoc_dataplane.dataplane_info
#   terraform state show data.ibm_byoc_engine.db2_engine_info
#
# UPDATE:
#   1. Modify variables in terraform.tfvars
#   2. terraform apply
#
# DELETE:
#   terraform destroy
#
# ============================================================================
# DATA SOURCE USAGE EXAMPLES
# ============================================================================
#
# Example 1: Query existing Db2 dataplane and engine
# ----------------------------------------------------
# terraform init
# terraform plan   # Shows what data will be fetched
# terraform apply  # Fetches the data
# terraform output dataplane_details  # Display dataplane info
# terraform output db2_engine_details # Display engine info
#
# Example 2: List all resources in subscription
# ----------------------------------------------
# terraform output all_dataplanes  # Shows all dataplanes
# terraform output all_engines     # Shows all engines on current dataplane
#
# Example 3: Use data in other resources
# ---------------------------------------
# You can reference data source outputs in other resources:
#   resource "some_resource" "example" {
#     dataplane_id = data.ibm_byoc_dataplane.dataplane_info.id
#     engine_endpoint = data.ibm_byoc_engine.db2_engine_info.endpoint
#   }
#
# ============================================================================