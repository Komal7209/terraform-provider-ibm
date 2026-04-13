# ============================================================================
# IBM BYOC Terraform Outputs
# ============================================================================
# These outputs display the results of CRUD operations
# ============================================================================

# ----------------------------------------------------------------------------
# Dataplane Outputs
# ----------------------------------------------------------------------------

output "dataplane_id" {
  description = "The ID of the created dataplane"
  value       = ibm_byoc_dataplane.byoc_dataplane.dataplane_id
}

output "dataplane_name" {
  description = "The name of the dataplane"
  value       = ibm_byoc_dataplane.byoc_dataplane.name
}

output "dataplane_status" {
  description = "The current status of the dataplane"
  value       = ibm_byoc_dataplane.byoc_dataplane.status
}

output "dataplane_region" {
  description = "The region where the dataplane is deployed"
  value       = ibm_byoc_dataplane.byoc_dataplane.region
}

output "dataplane_cloud_provider" {
  description = "The cloud provider for the dataplane"
  value       = ibm_byoc_dataplane.byoc_dataplane.cloud_provider
}

output "dataplane_created_at" {
  description = "Timestamp when the dataplane was created"
  value       = ibm_byoc_dataplane.byoc_dataplane.created_at
}

output "dataplane_updated_at" {
  description = "Timestamp when the dataplane was last updated"
  value       = ibm_byoc_dataplane.byoc_dataplane.updated_at
}

# ----------------------------------------------------------------------------
# Engine Outputs
# ----------------------------------------------------------------------------

output "db2_engine_id" {
  description = "The ID of the Db2 engine"
  value       = ibm_byoc_engine.db2_engine.engine_id
}

output "db2_engine_name" {
  description = "The name of the Db2 engine"
  value       = ibm_byoc_engine.db2_engine.engine_name
}

output "db2_engine_type" {
  description = "The type of the engine"
  value       = ibm_byoc_engine.db2_engine.engine_type
}

output "db2_engine_status" {
  description = "The current status of the Db2 engine"
  value       = ibm_byoc_engine.db2_engine.status
}

output "db2_engine_endpoint" {
  description = "The connection endpoint for the Db2 engine"
  value       = ibm_byoc_engine.db2_engine.endpoint
  sensitive   = true
}

output "db2_engine_created_at" {
  description = "Timestamp when the engine was created"
  value       = ibm_byoc_engine.db2_engine.created_at
}

output "db2_engine_updated_at" {
  description = "Timestamp when the engine was last updated"
  value       = ibm_byoc_engine.db2_engine.updated_at
}

# ----------------------------------------------------------------------------
# Data Source Outputs (for verification)
# ----------------------------------------------------------------------------

output "dataplane_info_from_datasource" {
  description = "Dataplane information retrieved via data source"
  value = {
    id             = data.ibm_byoc_dataplane.dataplane_info.dataplane_id
    name           = data.ibm_byoc_dataplane.dataplane_info.name
    status         = data.ibm_byoc_dataplane.dataplane_info.status
    cloud_provider = data.ibm_byoc_dataplane.dataplane_info.cloud_provider
    region         = data.ibm_byoc_dataplane.dataplane_info.region
  }
}

output "all_dataplanes_count" {
  description = "Total number of dataplanes in the subscription"
  value       = length(data.ibm_byoc_dataplanes.all_dataplanes.dataplanes)
}

output "db2_engine_info_from_datasource" {
  description = "Engine information retrieved via data source"
  value = {
    id          = data.ibm_byoc_engine.db2_engine_info.engine_id
    name        = data.ibm_byoc_engine.db2_engine_info.engine_name
    type        = data.ibm_byoc_engine.db2_engine_info.engine_type
    status      = data.ibm_byoc_engine.db2_engine_info.status
  }
}

output "all_engines_count" {
  description = "Total number of engines on the dataplane"
  value       = length(data.ibm_byoc_engines.all_engines.engines)
}

# ----------------------------------------------------------------------------
# Resource Configuration Summary
# ----------------------------------------------------------------------------

output "deployment_summary" {
  description = "Summary of the deployed resources"
  value = {
    subscription_id = var.subscription_id
    dataplane = {
      id             = ibm_byoc_dataplane.byoc_dataplane.dataplane_id
      name           = ibm_byoc_dataplane.byoc_dataplane.name
      cloud_provider = ibm_byoc_dataplane.byoc_dataplane.cloud_provider
      region         = ibm_byoc_dataplane.byoc_dataplane.region
      status         = ibm_byoc_dataplane.byoc_dataplane.status
    }
    engine = {
      id     = ibm_byoc_engine.db2_engine.engine_id
      name   = ibm_byoc_engine.db2_engine.engine_name
      type   = ibm_byoc_engine.db2_engine.engine_type
      status = ibm_byoc_engine.db2_engine.status
    }
  }
}

# ----------------------------------------------------------------------------
# CRUD Operation Verification
# ----------------------------------------------------------------------------

output "crud_verification" {
  description = "Verification that CRUD operations are working"
  value = {
    create_successful = ibm_byoc_dataplane.byoc_dataplane.status != "" && ibm_byoc_engine.db2_engine.status != ""
    read_successful   = data.ibm_byoc_dataplane.dataplane_info.dataplane_id == ibm_byoc_dataplane.byoc_dataplane.dataplane_id
    update_ready      = "Modify terraform.tfvars and run 'terraform apply' to test UPDATE"
    delete_ready      = "Run 'terraform destroy' to test DELETE"
  }
}

# ============================================================================
# Usage Examples
# ============================================================================
#
# View all outputs:
#   terraform output
#
# View specific output:
#   terraform output dataplane_id
#   terraform output deployment_summary
#
# View sensitive output:
#   terraform output -json db2_engine_endpoint
#
# ============================================================================