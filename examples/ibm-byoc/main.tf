// Provision byoc_engine resource instance
resource "ibm_byoc_engine" "byoc_engine_instance" {
  subscription_id = var.byoc_engine_subscription_id
  dataplane_id = var.byoc_engine_dataplane_id
  availability_zone = var.byoc_engine_availability_zone
  storage_units = var.byoc_engine_storage_units
  compute_units = var.byoc_engine_compute_units
  engine_name = var.byoc_engine_engine_name
  engine_type = var.byoc_engine_engine_type
  admin_username = var.byoc_engine_admin_username
  admin_password = var.byoc_engine_admin_password
  admin_email = var.byoc_engine_admin_email
  endpoint_type = var.byoc_engine_endpoint_type
  instance_type = var.byoc_engine_instance_type
  replicas = var.byoc_engine_replicas
  public_enabled = var.byoc_engine_public_enabled
  private_link_service_enabled = var.byoc_engine_private_link_service_enabled
  oracle_compatibility = var.byoc_engine_oracle_compatibility
  plan = var.byoc_engine_plan
  profile_name = var.byoc_engine_profile_name
  service_principals = var.byoc_engine_service_principals
  subscription_ids = var.byoc_engine_subscription_ids
}
