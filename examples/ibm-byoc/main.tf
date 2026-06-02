// ============================================================================
// Local Variables - Generate unique engine name with timestamp
// ============================================================================

locals {
  # Generate a unique engine name by appending timestamp
  # Format: db2-engine-20260602-143622
  engine_name_with_timestamp = "${var.byoc_engine_engine_name}-${formatdate("YYYYMMDDhhmmss", timestamp())}"
}

// ============================================================================
// BYOC Engine Resource - Create a new BYOC DB2 Engine
// ============================================================================

# resource "ibm_byoc_engine" "byoc_engine_instance" {
#   subscription_id               = var.byoc_engine_subscription_id
#   dataplane_id                  = var.byoc_engine_dataplane_id
#   storage_units                 = var.byoc_engine_storage_units
#   compute_units                 = var.byoc_engine_compute_units
#   engine_name                   = local.engine_name_with_timestamp
#   engine_type                   = var.byoc_engine_engine_type
#   admin_username                = var.byoc_engine_admin_username
#   admin_password                = var.byoc_engine_admin_password
#   admin_email                   = var.byoc_engine_admin_email
#   endpoint_type                 = var.byoc_engine_endpoint_type
#   instance_type                 = var.byoc_engine_instance_type
#   replicas                      = var.byoc_engine_replicas
#   public_enabled                = var.byoc_engine_public_enabled
#   private_link_service_enabled  = var.byoc_engine_private_link_service_enabled
#   oracle_compatibility          = var.byoc_engine_oracle_compatibility
# }

// ============================================================================
// Data Source - Get details of a specific BYOC Engine by ID
// ============================================================================

# Note: The engine_id needs to be extracted from the composite resource ID
# Resource ID format: subscription_id/dataplane_id/engine_id
# For now, commenting out until engine is created and ID is known

data "ibm_byoc_engine" "byoc_engine_data" {
  subscription_id = var.byoc_engine_subscription_id
  dataplane_id    = var.byoc_engine_dataplane_id
  engine_id       =var.byoc_engine_id 
  # split("/", ibm_byoc_engine.byoc_engine_instance.id)[2] #For using same engine_id
}

// ============================================================================
// Data Source - List all BYOC Engines in a Dataplane
// ============================================================================

# data "ibm_byoc_engines" "byoc_engines_list" {
#   subscription_id = var.byoc_engine_subscription_id
#   dataplane_id    = var.byoc_engine_dataplane_id
# }

// ============================================================================
// Example: Using DB2 User Management with BYOC Engine
// ============================================================================

# Uncomment the following to create users in the BYOC engine using existing DB2 SDK

# resource "ibm_db2_user" "byoc_admin_user" {
#   deployment_id = ibm_byoc_engine.byoc_engine_instance.engine_id
#
#   id       = "byoc-admin"
#   name     = "byoc_admin"
#   password = var.db2_user_password
#   role     = "bluadmin"
#   email    = "byoc-admin@example.com"
#   locked   = "no"
#
#   authentication {
#     method    = "internal"
#     policy_id = "Default"
#   }
# }

# resource "ibm_db2_user" "byoc_app_user" {
#   deployment_id = ibm_byoc_engine.byoc_engine_instance.engine_id
#
#   id       = "app-user"
#   name     = "application_user"
#   password = var.db2_user_password
#   role     = "bluuser"
#   email    = "app-user@example.com"
#   locked   = "no"
#
#   authentication {
#     method    = "internal"
#     policy_id = "Default"
#   }
#
#   depends_on = [ibm_db2_user.byoc_admin_user]
# }
